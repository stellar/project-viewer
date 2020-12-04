package queries

import (
	"context"

	"cloud.google.com/go/bigquery"
)

const queryLimit = 100

// CorridorResult is the result of a corridor query. It contains the source and destination volumes for a given ledger sequence
type CorridorResult struct {
	Seq    int64   `json:"ledgerSequence"`
	Source float64 `json:"sourceSum"`
	Dest   float64 `json:"destinationSum"`
}

// VolumeResult is the result of a volume to/from query. It contains the volume for a given ledger sequence
type VolumeResult struct {
	Seq    int64   `json:"ledgerSequence"`
	Volume float64 `json:"volume"`
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
