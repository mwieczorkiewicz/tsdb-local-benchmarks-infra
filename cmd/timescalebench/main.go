package main

import (
	"context"
	"log"
	"os"

	"github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/db/postgres"
	"github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/db/postgres/timescale"
	"github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/loader"
)

func main() {
	loadProfile := loader.LoadCsvDataLoadProfile("/Users/mikolajwieczorkiewicz/go/src/github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/data/subset/2021_ewz_bruttolastgang.csv")
	weather := loader.LoadCsvDataWeatherData("/Users/mikolajwieczorkiewicz/go/src/github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/data/subset/ugz_ogd_meteo_h1_2021.csv")

	conn, err := postgres.NewPostgresConnection()
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithCancel(context.Background())
	defer conn.Close(ctx)
	postgres.SetupTablesPostgres(ctx, conn, os.Getenv("SETUP_SCRIPT"))
	timescale.CreateHypertable(ctx, conn, "public.load_profile")
	timescale.CreateHypertable(ctx, conn, "public.weather_num")
	timescale.CreateHypertable(ctx, conn, "public.weather_string")

	postgres.InsertPostgresLoadProfileSingle(ctx, conn, loadProfile, "load-profile")
	postgres.InsertPostgresWeatherSingle(ctx, conn, weather, "weather-data")

}
