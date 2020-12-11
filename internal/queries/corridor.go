package queries

import (
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

// RunCorridorQuery queries BigQuery for the volume of assets over the specified corridor and returns the results
func RunCorridorQuery(source, dest Asset, startUnixTimestamp, endUnixTimestamp, aggregateBy string, client *bigquery.Client) ([]CorridorResult, error) {
	query := createCorridorQuery(source, dest, startUnixTimestamp, endUnixTimestamp, aggregateBy)
	it, err := runQuery(query, client)
	if err != nil {
		return nil, fmt.Errorf("error running query \n%s\n%v", query, err)
	}

	var results []CorridorResult
	for {
		var res CorridorResult
		if err := it.Next(&res); err == iterator.Done {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error parsing results from query: %v", err)
		}

		results = append(results, res)
	}

	return results, nil
}

func createCorridorQuery(source, dest Asset, startUnixTimestamp, endUnixTimestamp, aggregateBy string) string {
	paymentQuery := createCorridorPaymentQuery(source, dest, startUnixTimestamp, endUnixTimestamp, aggregateBy)
	tradeQuery := createCorridorTradeQuery(source, dest, startUnixTimestamp, endUnixTimestamp, aggregateBy)

	return createCombinedQuery(paymentQuery, tradeQuery)
}

// createCorridorTradeQuery returns a query that gets the total source volume through the corridor.
// The volume is calculated by looking at trades between the two assets within the timestamp range.
// The timestamps are in UTC to ensure they are consistent with the ledger closed_at timestamps.
func createCorridorTradeQuery(source, dest Asset, startUnixTimestamp, endUnixTimestamp, aggregateBy string) string {
	// A sample query is:
	// SELECT FORMAT("Ledger %d", L.sequence) AS title, SUM(T.base_amount)/10000000 as source, SUM(T.counter_amount)/10000000 as dest
	// FROM `crypto-stellar.crypto_stellar.history_trades` T
	// JOIN `crypto-stellar.crypto_stellar.history_assets` B ON B.id=T.base_asset_id
	// JOIN `crypto-stellar.crypto_stellar.history_assets` C ON C.id=T.counter_asset_id
	// JOIN `crypto-stellar.crypto_stellar.history_ledgers` L ON L.closed_at=T.ledger_closed_at
	// WHERE ((B.asset_code="NGNT" AND B.asset_issuer="GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD" AND C.asset_code="EURT"
	// AND C.asset_issuer="GAP5LETOV6YIE62YAM56STDANPRDO7ZFDBGSNHJQIYGGKSMOZAHOOS2S") OR
	// (B.asset_code="EURT" AND B.asset_issuer="GAP5LETOV6YIE62YAM56STDANPRDO7ZFDBGSNHJQIYGGKSMOZAHOOS2S" AND C.asset_code="NGNT"
	// AND C.asset_issuer="GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD"))
	// GROUP BY title ORDER BY title ASC LIMIT 100

	// For a given trade, if the source asset matches the trade's base asset, then we want the base amount of the trade
	baseAssetMatch := fmt.Sprintf("(B.asset_code=\"%s\" AND B.asset_issuer=\"%s\")", source.Code, source.Issuer)
	baseAssetSelect := "SUM(T.base_amount)/10000000"

	// For a given trade, if the source asset matches the trade's counter asset, then we want the counter amount of the trade
	counterAssetMatch := fmt.Sprintf("C.asset_code=\"%s\" AND C.asset_issuer=\"%s\"", source.Code, source.Issuer)
	counterAssetSelect := "SUM(T.counter_amount)/10000000"

	titleField := getTitleField("L.sequence", "closed_at", aggregateBy)
	query := fmt.Sprintf("SELECT %s, COUNT(history_operation_id) as count, CASE WHEN %s THEN %s WHEN %s THEN %s END AS amount ", titleField, baseAssetMatch, baseAssetSelect, counterAssetMatch, counterAssetSelect)
	query += " FROM `crypto-stellar.crypto_stellar.history_trades` T"
	query += " JOIN `crypto-stellar.crypto_stellar.history_assets` B ON B.id=T.base_asset_id"
	query += " JOIN `crypto-stellar.crypto_stellar.history_assets` C ON C.id=T.counter_asset_id"
	query += " JOIN `crypto-stellar.crypto_stellar.history_ledgers` L ON L.closed_at=T.ledger_closed_at"
	query += " WHERE " +
		fmt.Sprintf("((B.asset_code=\"%s\" AND B.asset_issuer=\"%s\" AND C.asset_code=\"%s\" AND C.asset_issuer=\"%s\")",
			source.Code, source.Issuer, dest.Code, dest.Issuer)
	query += " OR " +
		fmt.Sprintf("(B.asset_code=\"%s\" AND B.asset_issuer=\"%s\" AND C.asset_code=\"%s\" AND C.asset_issuer=\"%s\"))",
			dest.Code, dest.Issuer, source.Code, source.Issuer)
	if startUnixTimestamp != "" && endUnixTimestamp != "" {
		query += fmt.Sprintf(" AND closed_at BETWEEN TIMESTAMP_SECONDS(%s) AND TIMESTAMP_SECONDS(%s)", startUnixTimestamp, endUnixTimestamp)
	}

	query += fmt.Sprintf(" GROUP BY title, B.asset_code, B.asset_issuer, C.asset_code, C.asset_issuer ORDER BY title ASC LIMIT %d", queryLimit)
	return query
}

// createCorridorPaymentQuery returns a query that gets the total source through the corridor.
// The volume is calculated by looking at successful path payments that start from the source asset and end at the
// destination asset within the timestamp range. The timestamps are in UTC to ensure they are consistent with the ledger closed_at timestamps.
func createCorridorPaymentQuery(source, dest Asset, startUnixTimestamp, endUnixTimestamp, aggregateBy string) string {
	// A sample query is below:
	// SELECT FORMAT("Ledger %d", ledger_sequence) AS title, SUM(source_amount) AS source, SUM(amount) AS dest FROM
	// `crypto-stellar.crypto_stellar.enriched_history_operations` WHERE (type=2 OR type=13) AND successful=true
	// AND (source_asset_code="NGNT" AND source_asset_issuer="GAWODAROMJ33V5YDFY3NPYTHVYQG7MJXVJ2ND3AOGIHYRWINES6ACCPD" AND asset_code="EURT"
	// AND asset_issuer="GAP5LETOV6YIE62YAM56STDANPRDO7ZFDBGSNHJQIYGGKSMOZAHOOS2S")
	// GROUP BY title ORDER BY title ASC LIMIT 100

	titleField := getTitleField("ledger_sequence", "closed_at", aggregateBy)
	query := fmt.Sprintf("SELECT %s, COUNT(op_id) as count, SUM(source_amount) AS amount, FROM `crypto-stellar.crypto_stellar.enriched_history_operations`", titleField)
	query += " WHERE (type=2 OR type=13) AND successful=true"
	query += " AND " +
		fmt.Sprintf("(source_asset_code=\"%s\" AND source_asset_issuer=\"%s\" AND asset_code=\"%s\" AND asset_issuer=\"%s\")",
			source.Code, source.Issuer, dest.Code, dest.Issuer)
	if startUnixTimestamp != "" && endUnixTimestamp != "" {
		query += fmt.Sprintf(" AND closed_at BETWEEN TIMESTAMP_SECONDS(%s) AND TIMESTAMP_SECONDS(%s)", startUnixTimestamp, endUnixTimestamp)
	}

	query += fmt.Sprintf(" GROUP BY title ORDER BY title ASC LIMIT %d", queryLimit)
	return query
}
