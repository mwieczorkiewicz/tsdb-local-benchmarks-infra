package main

import (
	"context"
	"log"

	"github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/db/influx"
	"github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/loader"
)

func main() {
	loadProfile := loader.LoadCsvDataLoadProfile("/Users/mikolajwieczorkiewicz/go/src/github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/data/subset/2021_ewz_bruttolastgang.csv")
	weather := loader.LoadCsvDataWeatherData("/Users/mikolajwieczorkiewicz/go/src/github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/data/subset/ugz_ogd_meteo_h1_2021.csv")

	weatherPoints := influx.ConvertWeatherDataSetToPoints(weather)
	loadProfilePoints := influx.ConvertLoadProfileDataSetToPoints(loadProfile)

	client, err := influx.NewInfluxConnection(500)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	ctx := context.Background()
	cancelCtx, cancelFunc := context.WithCancel(ctx)

	influx.InsertPointsToInfluxSingle(cancelCtx, client, "iot-bucket", "acme", "weather-data", weatherPoints)
	influx.InsertPointsToInfluxBatch(cancelCtx, client, "iot-bucket", "acme", "weather-data", weatherPoints)

	influx.InsertPointsToInfluxSingle(cancelCtx, client, "iot-bucket", "acme", "load-prof-data", loadProfilePoints)
	influx.InsertPointsToInfluxBatch(cancelCtx, client, "iot-bucket", "acme", "load-prof-data", loadProfilePoints)

	cancelFunc()
}
