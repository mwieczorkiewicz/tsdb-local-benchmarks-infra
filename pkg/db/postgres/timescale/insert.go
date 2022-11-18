package timescale

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/util"
)

const dbName = "timescale"

func InsertPointsToInfluxSingle(ctx context.Context, c influxdb2.Client, bucket, org, dataSet string, data []*write.Point) error {
	defer util.TimeTrack(time.Now(), dataSet, dbName)
	writeAPI := c.WriteAPIBlocking(org, bucket)
	for _, p := range data {
		err := writeAPI.WritePoint(ctx, p)
		if err != nil {
			return err
		}
		err = writeAPI.Flush(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertPointsToInfluxBatch(ctx context.Context, c influxdb2.Client, bucket, org, dataSet string, data []*write.Point) error {
	defer util.TimeTrack(time.Now(), dataSet, dbName)
	writeAPI := c.WriteAPIBlocking(org, bucket)
	writeAPI.EnableBatching()
	for _, p := range data {
		err := writeAPI.WritePoint(ctx, p)
		if err != nil {
			return err
		}
	}
	return writeAPI.Flush(ctx)
}
