package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"

	"cloud.google.com/go/bigquery"
	"github.com/go-chi/chi"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

const queryLimit = 100
const projectID = "test-project-291320"
const keyFileName = "testingKey.json"

type queryResult struct {
	Seq    int64   `json:"ledgerSequence"`
	Source float64 `json:"sourceSum"`
	Dest   float64 `json:"destinationSum"`
}

// Asset represents an asset
type Asset struct {
	Code   string
	Issuer string
}

func (a Asset) isCompleteAsset() bool {
	return a.Code != "" && a.Issuer != ""
}

func createQuery(s, d Asset, start, end string) string {
	query := "SELECT ledger_sequence as seq, SUM(source_amount) AS source, SUM(amount) AS dest FROM `crypto-stellar.crypto_stellar.enriched_history_operations` WHERE (type=2 OR type=13) AND successful=true"
	query += " AND " +
		fmt.Sprintf("(source_asset_code=\"%s\" AND source_asset_issuer=\"%s\" AND asset_code=\"%s\" AND asset_issuer=\"%s\")",
			s.Code, s.Issuer, d.Code, d.Issuer)
	if start != "" && end != "" {
		query += " AND ledger_sequence BETWEEN " + start + " AND " + end
	}

	query += fmt.Sprintf(" GROUP BY ledger_sequence LIMIT %d", queryLimit)
	return query
}

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

func getQueryResults(it *bigquery.RowIterator) ([]queryResult, error) {
	var results []queryResult
	for {
		var res queryResult
		if err := it.Next(&res); err == iterator.Done {
			break
		} else if err != nil {
			return nil, err
		}

		results = append(results, res)
	}

	return results, nil
}

func queryHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		source := Asset{
			Code:   r.FormValue("sourceCode"),
			Issuer: r.FormValue("sourceIssuer"),
		}

		dest := Asset{
			Code:   r.FormValue("destCode"),
			Issuer: r.FormValue("destIssuer"),
		}

		startSeq := r.FormValue("start")
		endSeq := r.FormValue("end")
		if !source.isCompleteAsset() || !dest.isCompleteAsset() {
			http.Error(w, "Please connect to this URL with parameters sourceCode, sourceIssuer, destCode, destIssuer", 400)
			return
		}

		query := createQuery(source, dest, startSeq, endSeq)
		it, err := runQuery(query)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		results, err := getQueryResults(it)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		resultsMap := map[string][]queryResult{
			"results": results,
		}

		marshalled, err := json.Marshal(resultsMap)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		fmt.Fprintf(w, string(marshalled))
	})
}

// ServeMux creates Mux to serve api
func ServeMux() http.Handler {
	mux := chi.NewMux()
	mux.Handle("/", queryHandler())
	return mux
}

func main() {
	fmt.Println("Starting server on port :8080")
	mux := ServeMux()
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}