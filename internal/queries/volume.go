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

// createVolumeTradeQuery returns a query that gets the the total volume to/from an asset, grouped by ledger.
// The volume is calculated by looking at trades involving the assets within the timestamp range.
// The timestamps are in UTC to ensure they are consistent with the ledger closed_at timestamps.
func createVolumeTradeQuery(asset Asset, volumeFrom bool, startUnixTimestamp, endUnixTimestamp, aggregateBy string) string {
	// A sample query is below:
	// SELECT FORMAT("Ledger %d", L.sequence) AS title,
	// CASE WHEN (B.asset_code="NGNT" AND B.asset_issuer="GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD") THEN SUM(T.counter_amount)/10000000
	// WHEN C.asset_code="NGNT" AND C.asset_issuer="GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD" THEN SUM(T.base_amount)/10000000 END as volume,
	// FROM `crypto-stellar.crypto_stellar.history_trades` T
	// JOIN `crypto-stellar.crypto_stellar.history_assets` B ON B.id=T.base_asset_id
	// JOIN `crypto-stellar.crypto_stellar.history_assets` C ON C.id=T.counter_asset_id
	// JOIN `crypto-stellar.crypto_stellar.history_ledgers` L ON L.closed_at=T.ledger_closed_at
	// WHERE ((B.asset_code="NGNT" AND B.asset_issuer="GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD") OR
	// C.asset_code="NGNT" AND C.asset_issuer="GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD")
	// GROUP BY title, B.asset_code, B.asset_issuer, C.asset_code, C.asset_issuer
	// ORDER BY title ASC LIMIT 100

	// Construct the base match and select statements. If we want volume from the base asset,
	// we need to look at the base_amount. If we want volume to, we look at the counter_amount.
	assetType := "counter"
	if volumeFrom {
		assetType = "base"
	}

	baseAssetMatch := fmt.Sprintf("(B.asset_code=\"%s\" AND B.asset_issuer=\"%s\")", asset.Code, asset.Issuer)
	baseAssetSelect := fmt.Sprintf("SUM(T.%s_amount)/10000000", assetType)

	// Construct the counter match and select statements. If we want volume from the counter asset,
	// we need to look at the counter_amount. If we want volume to, we look at the base_amount.
	assetType = "base"
	if volumeFrom {
		assetType = "counter"
	}

	counterAssetMatch := fmt.Sprintf("C.asset_code=\"%s\" AND C.asset_issuer=\"%s\"", asset.Code, asset.Issuer)
	counterAssetSelect := fmt.Sprintf("SUM(T.%s_amount)/10000000", assetType)

	titleQuery := getTitleField("L.sequence", "L.closed_at", aggregateBy)

	query := fmt.Sprintf("SELECT %s, CASE WHEN %s THEN %s WHEN %s THEN %s END as volume,",
		titleQuery, baseAssetMatch, baseAssetSelect, counterAssetMatch, counterAssetSelect)
	query += " FROM `crypto-stellar.crypto_stellar.history_trades` T"
	query += " JOIN `crypto-stellar.crypto_stellar.history_assets` B ON B.id=T.base_asset_id"
	query += " JOIN `crypto-stellar.crypto_stellar.history_assets` C ON C.id=T.counter_asset_id"
	query += " JOIN `crypto-stellar.crypto_stellar.history_ledgers` L ON L.closed_at=T.ledger_closed_at"
	query += fmt.Sprintf(" WHERE (%s OR %s)", baseAssetMatch, counterAssetMatch)

	if startUnixTimestamp != "" && endUnixTimestamp != "" {
		query += fmt.Sprintf(" AND L.closed_at BETWEEN TIMESTAMP_SECONDS(%s) AND TIMESTAMP_SECONDS(%s)", startUnixTimestamp, endUnixTimestamp)
	}

	query += fmt.Sprintf(" GROUP BY title, B.asset_code, B.asset_issuer, C.asset_code, C.asset_issuer ORDER BY title ASC LIMIT %d", queryLimit)
	return query
}

// createVolumeQuery returns a query that gets the total volume to/from an asset, grouped by ledger.
// If volumeFrom is true, then we get the volume from the asset.
// The volume is calculated by looking at successful path payments involving the asset within the timestamp range.
// The timestamps are in UTC to ensure they are consistent with the ledger closed_at timestamps.
func createVolumeQuery(asset Asset, volumeFrom bool, startUnixTimestamp, endUnixTimestamp, aggregateBy string) string {
	// A sample query is below:
	// SELECT FORMAT("Ledger %d", ledger_sequence) AS title, SUM(amount) AS volume FROM `crypto-stellar.crypto_stellar.enriched_history_operations`
	// WHERE (type=2 OR type=13) AND successful=true AND (asset_code="NGNT" AND asset_issuer="GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD")
	// GROUP BY title ORDER BY title ASC LIMIT 100

	equalityPrefix := ""
	if volumeFrom {
		equalityPrefix = "source_"
	}

	titleQuery := getTitleField("ledger_sequence", "closed_at", aggregateBy)

	query := fmt.Sprintf("SELECT %s, SUM(%samount) AS volume", titleQuery, equalityPrefix)
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
