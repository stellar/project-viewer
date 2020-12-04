package backend

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"

	"cloud.google.com/go/bigquery"
	"github.com/stellar/project-viewer/internal/queries"
	"google.golang.org/api/option"
)

const projectID = "test-project-291320"
const keyFileName = "testingKey.json"

var client *bigquery.Client

func getBigQueryClient() (*bigquery.Client, error) {
	if client != nil {
		return client, nil
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	newClient, err := bigquery.NewClient(context.Background(), projectID, option.WithCredentialsFile(path.Join(currentDir, keyFileName)))
	if err != nil {
		return nil, err
	}

	client = newClient
	return client, nil
}

// CorridorHandler processes the source and destination assets, makes a BigQuery query, and returns the results
func CorridorHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		source := queries.Asset{
			Code:   r.FormValue("sourceCode"),
			Issuer: r.FormValue("sourceIssuer"),
		}

		dest := queries.Asset{
			Code:   r.FormValue("destCode"),
			Issuer: r.FormValue("destIssuer"),
		}

		startSeq := r.FormValue("start")
		endSeq := r.FormValue("end")
		if !source.IsCompleteAsset() || !dest.IsCompleteAsset() {
			http.Error(w, "Please connect to this URL with parameters sourceCode, sourceIssuer, destCode, destIssuer", 400)
			return
		}

		client, err := getBigQueryClient()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating BigQuery client: %s", err), 500)
		}

		results, err := queries.RunCorridorQuery(source, dest, startSeq, endSeq, client)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		resultsMap := map[string][]queries.CorridorResult{
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

// VolumeHandler processes asset, makes a BigQuery query for the volume to or from that asset, and returns the results
func VolumeHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		asset := queries.Asset{
			Code:   r.FormValue("code"),
			Issuer: r.FormValue("issuer"),
		}

		volumeFromParam := r.FormValue("volumeFrom")
		volumeFrom := volumeFromParam != ""

		startSeq := r.FormValue("start")
		endSeq := r.FormValue("end")
		if !asset.IsCompleteAsset() {
			http.Error(w, "Please connect to this URL with parameters code and issuer", 400)
			return
		}

		client, err := getBigQueryClient()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating BigQuery client: %s", err), 500)
		}

		results, err := queries.RunVolumeQuery(asset, volumeFrom, startSeq, endSeq, client)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		resultsMap := map[string][]queries.VolumeResult{
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
