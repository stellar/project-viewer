package queries

import (
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

// RunVolumeQuery queries BigQuery for the volume of assets over the specified corridor and returns the results
func RunVolumeQuery(asset Asset, volumeFrom bool, startUnixTimestamp, endUnixTimestamp, aggregateBy string, client *bigquery.Client) ([]VolumeResult, error) {
	query := createVolumeQuery(asset, volumeFrom, startUnixTimestamp, endUnixTimestamp, aggregateBy)
	it, err := runQuery(query, client)
	if err != nil {
		return nil, fmt.Errorf("error running query \n%s\n%v", query, err)
	}

	var results []VolumeResult
	for {
		var res VolumeResult
		if err := it.Next(&res); err == iterator.Done {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error parsing results from query: %v", err)
		}

		results = append(results, res)
	}

	return results, nil
}

// createVolumeQuery creates a query that combines the volume trade query and volume payment query
func createVolumeQuery(asset Asset, volumeFrom bool, startUnixTimestamp, endUnixTimestamp, aggregateBy string) string {
	paymentQuery := createVolumePaymentQuery(asset, volumeFrom, startUnixTimestamp, endUnixTimestamp, aggregateBy)
	tradeQuery := createVolumeTradeQuery(asset, volumeFrom, startUnixTimestamp, endUnixTimestamp, aggregateBy)

	return createCombinedQuery(paymentQuery, tradeQuery)
}

// createVolumeTradeQuery returns a query that gets the the total volume to/from an asset, grouped by ledger.
// The volume is calculated by looking at trades involving the assets within the timestamp range.
// The timestamps are in UTC to ensure they are consistent with the ledger closed_at timestamps.
func createVolumeTradeQuery(asset Asset, volumeFrom bool, startUnixTimestamp, endUnixTimestamp, aggregateBy string) string {
	// A sample query is below:
	// WITH base_trades AS (
	// 	 SELECT FORMAT_DATE("%Y/%m/%d", DATE_TRUNC(DATE(closed_at), MONTH)) AS title, COUNT(history_operation_id) AS count,
	// 		CASE WHEN (B.asset_code="NGNT" AND B.asset_issuer="GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD") THEN SUM(T.base_amount)/10000000 END as amount
	// 		FROM `crypto-stellar.crypto_stellar.history_trades` T
	// 		JOIN `crypto-stellar.crypto_stellar.history_assets` B ON B.id=T.base_asset_id
	// 		JOIN `crypto-stellar.crypto_stellar.history_assets` C ON C.id=T.counter_asset_id
	// 		JOIN `crypto-stellar.crypto_stellar.history_ledgers` L ON L.closed_at=T.ledger_closed_at
	// 		WHERE (B.asset_code="NGNT" AND B.asset_issuer="GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD")
	// 		GROUP BY title, B.asset_code, B.asset_issuer
	// 		ORDER BY title ASC
	//   ),
	// 	counter_trades AS (
	// 	  SELECT FORMAT_DATE("%Y/%m/%d", DATE_TRUNC(DATE(closed_at), MONTH)) AS title, COUNT(history_operation_id) AS count,
	// 	    CASE WHEN (C.asset_code="NGNT" AND C.asset_issuer="GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD") THEN SUM(T.counter_amount)/10000000 END AS amount,
	// 		FROM `crypto-stellar.crypto_stellar.history_trades` T
	// 		JOIN `crypto-stellar.crypto_stellar.history_assets` B ON B.id=T.base_asset_id
	// 		JOIN `crypto-stellar.crypto_stellar.history_assets` C ON C.id=T.counter_asset_id
	// 		JOIN `crypto-stellar.crypto_stellar.history_ledgers` L ON L.closed_at=T.ledger_closed_at
	// 		WHERE (C.asset_code="NGNT" AND C.asset_issuer="GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD")
	// 		GROUP BY title, C.asset_code, C.asset_issuer
	// 		ORDER BY title ASC)
	// 	SELECT IFNULL(bt.title, ct.title) AS title, IFNULL(bt.count, 0) + IFNULL(ct.count, 0) as count, IFNULL(bt.amount, 0) + IFNULL(ct.amount, 0) as amount
	// 	FROM base_trades bt
	// 	FULL JOIN counter_trades ct ON bt.title=ct.title ORDER BY title ASC LIMIT 100

	titleQuery := getTitleField("L.sequence", "L.closed_at", aggregateBy)
	joins := " FROM `crypto-stellar.crypto_stellar.history_trades` T"
	joins += " JOIN `crypto-stellar.crypto_stellar.history_assets` B ON B.id=T.base_asset_id"
	joins += " JOIN `crypto-stellar.crypto_stellar.history_assets` C ON C.id=T.counter_asset_id"
	joins += " JOIN `crypto-stellar.crypto_stellar.history_ledgers` L ON L.closed_at=T.ledger_closed_at"
	// Construct the base match and select statements. If we want volume from the base asset,
	// we need to look at the base_amount. If we want volume to, we look at the counter_amount.
	assetType := "counter"
	if volumeFrom {
		assetType = "base"
	}

	baseAssetMatch := fmt.Sprintf("B.asset_code=\"%s\" AND B.asset_issuer=\"%s\"", asset.Code, asset.Issuer)
	baseAssetSelect := fmt.Sprintf("SUM(T.%s_amount)/10000000", assetType)
	baseAssetQuery := fmt.Sprintf("SELECT %s, COUNT(history_operation_id) AS count, CASE WHEN %s THEN %s END AS amount ", titleQuery, baseAssetMatch, baseAssetSelect)
	baseAssetQuery += joins
	baseAssetQuery += fmt.Sprintf(" WHERE (%s)", baseAssetMatch)
	if startUnixTimestamp != "" && endUnixTimestamp != "" {
		baseAssetQuery += fmt.Sprintf(" AND L.closed_at BETWEEN TIMESTAMP_SECONDS(%s) AND TIMESTAMP_SECONDS(%s)", startUnixTimestamp, endUnixTimestamp)
	}

	baseAssetQuery += "GROUP BY title, B.asset_code, B.asset_issuer ORDER BY title ASC"

	// Construct the counter match and select statements. If we want volume from the counter asset,
	// we need to look at the counter_amount. If we want volume to, we look at the base_amount.
	assetType = "base"
	if volumeFrom {
		assetType = "counter"
	}

	counterAssetMatch := fmt.Sprintf("C.asset_code=\"%s\" AND C.asset_issuer=\"%s\"", asset.Code, asset.Issuer)
	counterAssetSelect := fmt.Sprintf("SUM(T.%s_amount)/10000000", assetType)
	counterAssetQuery := fmt.Sprintf("SELECT %s, COUNT(history_operation_id) AS count, CASE WHEN %s THEN %s END AS amount ", titleQuery, counterAssetMatch, counterAssetSelect)
	counterAssetQuery += joins
	counterAssetQuery += fmt.Sprintf(" WHERE (%s)", counterAssetMatch)
	if startUnixTimestamp != "" && endUnixTimestamp != "" {
		counterAssetQuery += fmt.Sprintf(" AND L.closed_at BETWEEN TIMESTAMP_SECONDS(%s) AND TIMESTAMP_SECONDS(%s)", startUnixTimestamp, endUnixTimestamp)
	}

	counterAssetQuery += "GROUP BY title, C.asset_code, C.asset_issuer ORDER BY title ASC"

	query := fmt.Sprintf("WITH base_trades AS (%s), counter_trades AS (%s)", baseAssetQuery, counterAssetQuery)
	query += " SELECT IFNULL(bt.title, ct.title) AS title, IFNULL(bt.count, 0) + IFNULL(ct.count, 0) as count, IFNULL(bt.amount, 0) + IFNULL(ct.amount, 0) AS amount"
	query += " FROM base_trades bt"
	query += " FULL JOIN counter_trades ct "
	query += fmt.Sprintf(" ON bt.title=ct.title ORDER BY title ASC LIMIT %d", queryLimit)
	return query
}

// createVolumePaymentQuery returns a query that gets the total volume to/from an asset.
// If volumeFrom is true, then we get the volume from the asset.
// The volume is calculated by looking at successful path payments involving the asset within the timestamp range.
// The timestamps are in UTC to ensure they are consistent with the ledger closed_at timestamps.
func createVolumePaymentQuery(asset Asset, volumeFrom bool, startUnixTimestamp, endUnixTimestamp, aggregateBy string) string {
	// A sample query is below:
	// SELECT FORMAT("Ledger %d", ledger_sequence) AS title, COUNT(op_id) AS count, SUM(amount) AS amount
	// FROM `crypto-stellar.crypto_stellar.enriched_history_operations`
	// WHERE (type=2 OR type=13) AND successful=true AND (asset_code="NGNT" AND asset_issuer="GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD")
	// GROUP BY title ORDER BY title ASC LIMIT 100

	equalityPrefix := ""
	if volumeFrom {
		equalityPrefix = "source_"
	}

	titleQuery := getTitleField("ledger_sequence", "closed_at", aggregateBy)

	query := fmt.Sprintf("SELECT %s, COUNT(op_id) AS count, SUM(%samount) AS amount", titleQuery, equalityPrefix)
	query += " FROM `crypto-stellar.crypto_stellar.enriched_history_operations` WHERE (type=2 OR type=13) AND successful=true"
	query += " AND " +
		fmt.Sprintf("(%sasset_code=\"%s\" AND %sasset_issuer=\"%s\")",
			equalityPrefix, asset.Code, equalityPrefix, asset.Issuer)
	if startUnixTimestamp != "" && endUnixTimestamp != "" {
		query += fmt.Sprintf(" AND closed_at BETWEEN TIMESTAMP_SECONDS(%s) AND TIMESTAMP_SECONDS(%s)", startUnixTimestamp, endUnixTimestamp)
	}

	query += fmt.Sprintf(" GROUP BY title ORDER BY title ASC LIMIT %d", queryLimit)
	return query
}
