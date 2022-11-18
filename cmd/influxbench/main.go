package main

import "github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/loader"

func main() {
	loader.LoadCsvDataLoadProfile("/Users/mikolajwieczorkiewicz/go/src/github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/data/2021_ewz_bruttolastgang.csv")
	loader.LoadCsvDataWeatherData("/Users/mikolajwieczorkiewicz/go/src/github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/data/ugz_ogd_meteo_h1_2021.csv")
}
