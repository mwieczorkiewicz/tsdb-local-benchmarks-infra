package influx

import (
	"context"
	"errors"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/model"
	"github.com/shopspring/decimal"
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

func InsertLoadProfileDataToMeasurementSingle(ctx context.Context, c influxdb2.Client, bucket, org, measurement string, data []model.LoadProfileRecord) error {
	writeAPI := c.WriteAPIBlocking(org, bucket)
	for _, v := range data {
		f, _ := v.LoadProfile.Float64()
		p := influxdb2.NewPoint("load-profile",
			map[string]string{"status": v.Status},
			map[string]interface{}{"value": f},
			v.Ts.Ts)
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

func InsertLoadProfileDataToMeasurementBatch(ctx context.Context, c influxdb2.Client, bucket, org, measurement string, data []model.LoadProfileRecord) error {
	writeAPI := c.WriteAPIBlocking(org, bucket)
	writeAPI.EnableBatching()
	for _, v := range data {
		f, _ := v.LoadProfile.Float64()
		p := influxdb2.NewPoint("load-profile",
			map[string]string{"status": v.Status},
			map[string]interface{}{"value": f},
			v.Ts.Ts)
		writeAPI.WritePoint(ctx, p)
	}
	writeAPI.Flush(ctx)
	return nil
}

func InsertWeatherDataToMeasurementSingle(ctx context.Context, c influxdb2.Client, bucket, org, measurement string, data []model.WeatherRecord) error {
	writeAPI := c.WriteAPIBlocking(org, bucket)
	for _, v := range data {
		num, err := decimal.NewFromString(v.Value)
		if err != nil {
			f, _ := num.Float64()
			p := influxdb2.NewPoint("weather-data-num",
				map[string]string{
					"location":  v.Location,
					"parameter": v.Parameter,
					"unit":      v.Unit,
				},
				map[string]interface{}{"value": f},
				v.Ts.Ts)
			writeAPI.WritePoint(ctx, p)
		}
		writeAPI.Flush(ctx)
	}
	return nil
}
