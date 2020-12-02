package projectviewer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stellar/project-viewer/internal/queries"
)

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

		results, err := queries.RunCorridorQuery(source, dest, startSeq, endSeq)
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
