package queries

import (
	"fmt"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

// RunRateQuery queries BigQuery for the volume of assets over the specified corridor and returns the results
func RunRateQuery(source, dest Asset, startLedger, endLedger string, client *bigquery.Client) ([]RateResult, error) {
	query := createRateQuery(source, dest, startLedger, endLedger)
	it, err := runQuery(query, client)
	if err != nil {
		return nil, fmt.Errorf("error running query \n%s\n%v", query, err)
	}

	var results []RateResult
	for {
		var res RateResult
		if err := it.Next(&res); err == iterator.Done {
			break
		} else if err != nil {
			return nil, fmt.Errorf("error parsing results from query: %v", err)
		}

		results = append(results, res)
	}

	return results, nil
}

// createRateTradeQuery returns a query that gets the the rate between two assets, grouped by ledger.
// The volume is calculated by looking at trades involving the assets within the provided ledger range.
func createRateTradeQuery(source, dest Asset, startLedger, endLedger string) string {
	// If the assets map as we expect (source -> base and dest -> counter), then the rate
	// is the counter amount over the base amount. The rate convert from X source assets to Y dest assets
	// so the units for the rate should be (dest/source = counter/base)
	baseAssetMatch := fmt.Sprintf("((B.asset_code=\"%s\" AND B.asset_issuer=\"%s\") OR (C.asset_code=\"%s\" AND C.asset_issuer=\"%s\"))",
		source.Code, source.Issuer, dest.Code, dest.Issuer)
	baseAssetSelect := "SUM(T.counter_amount)/SUM(T.base_amount)"

	// If the assets map as the opposite of what we expect (source -> counter and dest -> base), then the rate
	// is the base amount over the counter amount. The rate convert from X source assets to Y dest assets
	// so the units for the rate should be (dest/source = base/counter)
	counterAssetMatch := fmt.Sprintf("((C.asset_code=\"%s\" AND C.asset_issuer=\"%s\") OR (B.asset_code=\"%s\" AND B.asset_issuer=\"%s\"))",
		source.Code, source.Issuer, dest.Code, dest.Issuer)
	counterAssetSelect := "SUM(T.base_amount)/SUM(T.counter_amount)"

	query := fmt.Sprintf("SELECT L.sequence AS seq, CASE WHEN %s THEN %s WHEN %s THEN %s END as rate,",
		baseAssetMatch, baseAssetSelect, counterAssetMatch, counterAssetSelect)
	query += " FROM `crypto-stellar.crypto_stellar.history_trades` T"
	query += " JOIN `crypto-stellar.crypto_stellar.history_assets` B ON B.id=T.base_asset_id"
	query += " JOIN `crypto-stellar.crypto_stellar.history_assets` C ON C.id=T.counter_asset_id"
	query += " JOIN `crypto-stellar.crypto_stellar.history_ledgers` L ON L.closed_at=T.ledger_closed_at"
	query += fmt.Sprintf(" WHERE (%s OR %s)", baseAssetMatch, counterAssetMatch)

	if startLedger != "" && endLedger != "" {
		query += fmt.Sprintf(" AND L.sequence BETWEEN %s AND %s", startLedger, endLedger)
	}

	query += fmt.Sprintf(" GROUP BY seq, B.asset_code, B.asset_issuer, C.asset_code, C.asset_issuer ORDER BY seq ASC LIMIT %d", queryLimit)
	return query
}

// createRateQuery returns a query that gets the on-DEX rate between two assets, grouped by ledger.
// The rate is calculated by looking at historical orderbooks. The average price of the highest bid
// and the lowest ask are averaged to get the rate at each ledger.
func createRateQuery(source, dest Asset, startLedger, endLedger string) string {
	normalMatch := fmt.Sprintf("(M.base_code=\"%s\" AND M.base_issuer=\"%s\" AND M.counter_code=\"%s\" AND M.counter_issuer=\"%s\")",
		source.Code, source.Issuer, dest.Code, dest.Issuer)
	reverseMatch := fmt.Sprintf("(M.base_code=\"%s\" AND M.base_issuer=\"%s\" AND M.counter_code=\"%s\" AND M.counter_issuer=\"%s\")",
		dest.Code, dest.Issuer, source.Code, source.Issuer)

	query := "WITH orderbooks AS ("
	query += " SELECT E.ledger_id, M.base_code, M.base_issuer, M.counter_code, M.counter_issuer,"
	query += ` ARRAY_AGG(CASE WHEN O.action="b" THEN O.price END IGNORE NULLS ORDER BY O.price DESC) AS bidPrices,`
	query += ` ARRAY_AGG(CASE WHEN O.action="s" THEN O.price END IGNORE NULLS ORDER BY O.price ASC) AS askPrices,`
	query += " FROM `hubble-261722.liquidity_data.fact_offer_events` AS E"
	query += " INNER JOIN `hubble-261722.liquidity_data.dim_offers` O ON (E.offer_instance_id = O.dim_offer_id)"
	query += " INNER JOIN `hubble-261722.liquidity_data.dim_markets` M ON (O.market_id = M.market_id)"
	query += fmt.Sprintf(" WHERE (%s OR %s)", normalMatch, reverseMatch)

	if startLedger != "" && endLedger != "" {
		query += fmt.Sprintf(" AND E.ledger_id BETWEEN %s AND %s", startLedger, endLedger)
	}

	query += " GROUP by E.ledger_id, M.base_code, M.base_issuer, M.counter_code, M.counter_issuer)"

	rateCalculation := "(orderbooks.askPrices[OFFSET(0)]+orderbooks.bidPrices[OFFSET(0)])/2"
	baseIsSource := fmt.Sprintf("orderbooks.base_code=\"%s\" AND orderbooks.base_issuer=\"%s\"", source.Code, source.Issuer)
	
	// if the base is not the source asset, then our rate is the reversed direction and so we must take the reciprocal
	query += fmt.Sprintf(" SELECT orderbooks.ledger_id as seq, CASE WHEN %s THEN %s ELSE 1/(%s) END as rate FROM orderbooks", baseIsSource, rateCalculation, rateCalculation)
	query += fmt.Sprintf(" WHERE %s IS NOT NULL", rateCalculation)
	query += fmt.Sprintf(" ORDER BY seq ASC LIMIT %d", queryLimit)
	return query
}
