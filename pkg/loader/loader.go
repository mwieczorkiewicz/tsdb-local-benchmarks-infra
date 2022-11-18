package loader

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/model"
)

const sep = ','

func LoadCsvDataLoadProfile(path string) []model.LoadProfileRecord {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	d := make([]model.LoadProfileRecord, 0)
	genericLoadData(f, &d)
	return d
}

func LoadCsvDataWeatherData(path string) []model.WeatherRecord {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	d := make([]model.WeatherRecord, 0)
	genericLoadData(f, &d)
	return d
}

func genericLoadData[T any](r io.Reader, ptr *T) error {
	reader := csv.NewReader(r)
	reader.Comma = sep
	reader.LazyQuotes = true
	err := gocsv.UnmarshalCSV(reader, ptr)
	if err != nil {
		return err
	}
	return nil
}
