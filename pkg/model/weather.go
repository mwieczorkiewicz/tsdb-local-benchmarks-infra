package model

import (
	"time"
)

type DateTimeWeather struct {
	Ts time.Time
}

type WeatherRecord struct {
	Ts        DateTimeWeather `csv:"Datum"`
	Location  string          `csv:"Standort"`
	Parameter string          `csv:"Parameter"`
	Value     string          `csv:"Wert"`
	Unit      string          `csv:"Einheit"`
}

// Convert the CSV string as internal date
func (dt *DateTimeWeather) UnmarshalCSV(csv string) (err error) {
	dt.Ts, err = time.Parse(tsTemperatureFormat, csv)
	return err
}
