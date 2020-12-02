package queries

import (
	"context"
	"os"
	"path"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

const queryLimit = 100
const projectID = "test-project-291320"
const keyFileName = "testingKey.json"

// CorridorResult is the result of a corridor query. It contains the source and destination volumes for a given ledger sequence
type CorridorResult struct {
	Seq    int64   `json:"ledgerSequence"`
	Source float64 `json:"sourceSum"`
	Dest   float64 `json:"destinationSum"`
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
func runQuery(query string) (*bigquery.RowIterator, error) {
	ctx := context.Background()

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	client, err := bigquery.NewClient(ctx, projectID, option.WithCredentialsFile(path.Join(currentDir, keyFileName)))
	if err != nil {
		return nil, err
	}

	q := client.Query(query)
	return q.Read(ctx)
}
