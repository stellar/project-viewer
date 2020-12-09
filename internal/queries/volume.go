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
	//	Construct the base match and select statements. If we want volume from the base asset,
	//	we need to look at the base_amount. If we want volume to, we look at the counter_amount.
	assetType := "counter"
	if volumeFrom {
		assetType = "base"
	}
	baseAssetMatch := fmt.Sprintf("(B.asset_code=\"%s\" AND B.asset_issuer=\"%s\")", asset.Code, asset.Issuer)
	baseAssetSelect := fmt.Sprintf("SUM(T.%s_amount)/10000000", assetType)

	//	Construct the counter match and select statements. If we want volume from the counter asset,
	//	we need to look at the counter_amount. If we want volume to, we look at the base_amount.
	assetType = "base"
	if volumeFrom {
		assetType = "counter"
	}
	counterAssetMatch := fmt.Sprintf("C.asset_code=\"%s\" AND C.asset_issuer=\"%s\"", asset.Code, asset.Issuer)
	counterAssetSelect := fmt.Sprintf("SUM(T.%s_amount)/10000000", assetType)

	query := "SELECT FORMAT(\"Ledger %d\", L.sequence) AS title," + fmt.Sprintf(" CASE WHEN %s THEN %s WHEN %s THEN %s END as volume,",
		baseAssetMatch, baseAssetSelect, counterAssetMatch, counterAssetSelect)
	query += " FROM `crypto-stellar.crypto_stellar.history_trades` T"
	query += " JOIN `crypto-stellar.crypto_stellar.history_assets` B ON B.id=T.base_asset_id"
	query += " JOIN `crypto-stellar.crypto_stellar.history_assets` C ON C.id=T.counter_asset_id"
	query += " JOIN `crypto-stellar.crypto_stellar.history_ledgers` L ON L.closed_at=T.ledger_closed_at"
	query += fmt.Sprintf(" WHERE (%s OR %s)", baseAssetMatch, counterAssetMatch)

	if startUnixTimestamp != "" && endUnixTimestamp != "" {
		query += fmt.Sprintf(" AND L.closed_at BETWEEN TIMESTAMP_SECONDS(%s) AND TIMESTAMP_SECONDS(%s)", startUnixTimestamp, endUnixTimestamp)
	}

	query += fmt.Sprintf(" GROUP BY title, B.asset_code, B.asset_issuer, C.asset_code, C.asset_issuer ORDER BY L.sequence ASC LIMIT %d", queryLimit)
	return query
}

// createVolumeQuery returns a query that gets the total volume to/from an asset, grouped by ledger.
// If volumeFrom is true, then we get the volume from the asset.
// The volume is calculated by looking at successful path payments involving the asset within the timestamp range.
// The timestamps are in UTC to ensure they are consistent with the ledger closed_at timestamps.
func createVolumeQuery(asset Asset, volumeFrom bool, startUnixTimestamp, endUnixTimestamp, aggregateBy string) string {
	equalityPrefix := ""
	if volumeFrom {
		equalityPrefix = "source_"
	}

	query := "SELECT FORMAT(\"Ledger %d\", ledger_sequence) AS title," + fmt.Sprintf(" SUM(%samount) AS volume", equalityPrefix)
	query += " FROM `crypto-stellar.crypto_stellar.enriched_history_operations` WHERE (type=2 OR type=13) AND successful=true"
	query += " AND " +
		fmt.Sprintf("(%sasset_code=\"%s\" AND %sasset_issuer=\"%s\")",
			equalityPrefix, asset.Code, equalityPrefix, asset.Issuer)
	if startUnixTimestamp != "" && endUnixTimestamp != "" {
		query += fmt.Sprintf(" AND closed_at BETWEEN TIMESTAMP_SECONDS(%s) AND TIMESTAMP_SECONDS(%s)", startUnixTimestamp, endUnixTimestamp)
	}

	query += fmt.Sprintf(" GROUP BY title ORDER BY ledger_sequence ASC LIMIT %d", queryLimit)
	return query
}
