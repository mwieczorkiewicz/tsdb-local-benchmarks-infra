package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type DateTimeLoadProfile struct {
	Ts time.Time
}

type LoadProfileRecord struct {
	Ts          DateTimeLoadProfile `csv:"zeitpunkt"`
	LoadProfile decimal.Decimal     `csv:"bruttolastgang"`
	Status      string              `csv:"status"`
}

// Convert the CSV string as internal date
func (dt *DateTimeLoadProfile) UnmarshalCSV(csv string) (err error) {
	dt.Ts, err = time.Parse(tsHourlyAggrFormat, csv)
	return err
}
