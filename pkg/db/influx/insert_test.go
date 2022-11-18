package influx

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/loader"
)

var numberOfIterations int = 1

func BenchmarkInsertWeatherPointsToInfluxSingle(b *testing.B) {
	b.N = numberOfIterations
	os.Setenv("INFLUXDB_URL", "http://192.168.0.200:8086")
	os.Setenv("INFLUXDB_TOKEN", "SecretToken")
	weather := loader.LoadCsvDataWeatherData("/Users/mikolajwieczorkiewicz/go/src/github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/data/subset/ugz_ogd_meteo_h1_2021.csv")

	weatherPoints := ConvertWeatherDataSetToPoints(weather)

	client, err := NewInfluxConnection(500)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)

	InsertPointsToInfluxSingle(cancelCtx, client, "iot-bucket", "acme", "weather-data", weatherPoints)
	cancelFunc()
}

func BenchmarkInsertWeatherPointsToInfluxBatch(b *testing.B) {
	b.N = 1
	os.Setenv("INFLUXDB_URL", "http://192.168.0.200:8086")
	os.Setenv("INFLUXDB_TOKEN", "SecretToken")
	weatherData := loader.LoadCsvDataWeatherData("/Users/mikolajwieczorkiewicz/go/src/github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/data/subset/ugz_ogd_meteo_h1_2021.csv")

	weatherPoints := ConvertWeatherDataSetToPoints(weatherData)

	client, err := NewInfluxConnection(500)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)

	InsertPointsToInfluxBatch(cancelCtx, client, "iot-bucket", "acme", "weather-data", weatherPoints)
	cancelFunc()
}

func BenchmarkInsertLoadPointsToInfluxSingle(b *testing.B) {
	b.N = 1
	os.Setenv("INFLUXDB_URL", "http://192.168.0.200:8086")
	os.Setenv("INFLUXDB_TOKEN", "SecretToken")
	loadProfile := loader.LoadCsvDataLoadProfile("/Users/mikolajwieczorkiewicz/go/src/github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/data/subset/2021_ewz_bruttolastgang.csv")

	loadProfPoints := ConvertLoadProfileDataSetToPoints(loadProfile)

	client, err := NewInfluxConnection(500)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)

	InsertPointsToInfluxSingle(cancelCtx, client, "iot-bucket", "acme", "weather-data", loadProfPoints)
	cancelFunc()
}

func BenchmarkInsertLoadPointsToInfluxBatch(b *testing.B) {
	b.N = 1
	os.Setenv("INFLUXDB_URL", "http://192.168.0.200:8086")
	os.Setenv("INFLUXDB_TOKEN", "SecretToken")
	loadProfile := loader.LoadCsvDataLoadProfile("/Users/mikolajwieczorkiewicz/go/src/github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/data/subset/2021_ewz_bruttolastgang.csv")

	loadProfPoints := ConvertLoadProfileDataSetToPoints(loadProfile)

	client, err := NewInfluxConnection(500)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)

	InsertPointsToInfluxBatch(cancelCtx, client, "iot-bucket", "acme", "weather-data", loadProfPoints)
	cancelFunc()
}
