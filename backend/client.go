package backend

import (
	"context"
	"os"
	"path"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/option"
)

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
