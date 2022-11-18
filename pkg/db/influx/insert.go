package influx

import (
	"context"
	"errors"
	"os"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/model"
	"github.com/shopspring/decimal"
)

func NewInfluxConnection() (influxdb2.Client, error) {
	dbToken := os.Getenv("INFLUXDB_TOKEN")
	if dbToken == "" {
		return nil, errors.New("auth token must be set - set INFLUXDB_TOKEN")
	}

	dbURL := os.Getenv("INFLUXDB_URL")
	if dbURL == "" {
		return nil, errors.New("url must be set - set INFLUXDB_URL")
	}

	// connect and check conn health
	client := influxdb2.NewClient(dbURL, dbToken)
	_, err := client.Health(context.Background())

	return client, err
}

func InsertLoadProfileDataToMeasurementSingle(c influxdb2.Client, bucket, org, measurement string, data []model.LoadProfileRecord) error {
	writeAPI := c.WriteAPI(org, bucket)
	for _, v := range data {
		f, _ := v.LoadProfile.Float64()
		p := influxdb2.NewPoint("load-profile",
			map[string]string{"status": v.Status},
			map[string]interface{}{"value": f},
			v.Ts.Ts)
		writeAPI.WritePoint(p)
		// Flush writes one by one
		writeAPI.Flush()
	}
	return nil
}

func InsertWeatherDataToMeasurementSingle(c influxdb2.Client, bucket, org, measurement string, data []model.WeatherRecord) error {
	writeAPI := c.WriteAPI(org, bucket)
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
			writeAPI.WritePoint(p)
		}
		writeAPI.Flush()
	}
	return nil
}
