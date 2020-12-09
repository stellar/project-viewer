package queries

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/bigquery"
)

const queryLimit = 100

// CorridorResult is the result of a corridor query. It contains the source and destination volumes for a given ledger sequence
type CorridorResult struct {
	Title  string  `json:"title"`
	Source float64 `json:"sourceSum"`
	Dest   float64 `json:"destinationSum"`
}

// VolumeResult is the result of a volume to/from query. It contains the volume for a given ledger sequence
type VolumeResult struct {
	Title  string  `json:"title"`
	Volume float64 `json:"volume"`
}

// RateResult is the result of a rate query. It contains the rate between two assets for a given ledger sequence
type RateResult struct {
	Title string  `json:"title"`
	Rate  float64 `json:"rate"`
}

// Asset represents an Asset with a code and Issuer
type Asset struct {
	Code   string
	Issuer string
}

// IsCompleteAsset returns true if the Asset has both a code and an issuer, and returns false otherwise
func (a Asset) IsCompleteAsset() bool {
	return a.Code != "" && a.Issuer != ""
}

// runQuery runs the provided query and returns the results
func runQuery(query string, client *bigquery.Client) (*bigquery.RowIterator, error) {
	ctx := context.Background()
	q := client.Query(query)
	return q.Read(ctx)
}

func getTitleField(ledgerID, closedAt, aggregateBy string) string {
	switch strings.ToLower(aggregateBy) {
	case "day":
		return fmt.Sprintf("DATE(%s) as title", closedAt)
	case "week", "month", "quarter", "year":
		return fmt.Sprintf("DATE_TRUNC(DATE(closed_at), %s) as title", strings.ToUpper(aggregateBy))
	default:
		return fmt.Sprintf("FORMAT(\"Ledger %%d\", %s) AS title", ledgerID)
	}
}
