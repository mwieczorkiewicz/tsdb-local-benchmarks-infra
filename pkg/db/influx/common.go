package influx

import (
	"context"
	"errors"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func NewInfluxConnection(batchSize uint) (influxdb2.Client, error) {
	dbToken := os.Getenv("INFLUXDB_TOKEN")
	if dbToken == "" {
		return nil, errors.New("auth token must be set - set INFLUXDB_TOKEN")
	}

	dbURL := os.Getenv("INFLUXDB_URL")
	if dbURL == "" {
		return nil, errors.New("url must be set - set INFLUXDB_URL")
	}

	// connect and check conn health
	client := influxdb2.NewClientWithOptions(dbURL, dbToken, influxdb2.DefaultOptions().SetBatchSize(batchSize))
	_, err := client.Health(context.Background())

	return client, err
}
