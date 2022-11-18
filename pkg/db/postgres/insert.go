package postgres

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/model"
	"github.com/mwieczorkiewicz/tsdb-local-benchmarks-infra/pkg/util"
	"github.com/shopspring/decimal"
)

const dbName = "psql"

const queryInsertTimeseriesLoadProfileData = `
INSERT INTO public.load_profile (ts, load_profile, status) VALUES ($1, $2, $3);
`

const queryInsertTimeseriesWeatherDataNum = `
INSERT INTO public.weather_num (ts, location, parameter, value, unit) VALUES ($1, $2, $3, $4, $5);
`

const queryInsertTimeseriesWeatherDataString = `
INSERT INTO public.weather_string (ts, location, parameter, value, unit) VALUES ($1, $2, $3, $4, $5);
`

func InsertPostgresLoadProfileSingle(ctx context.Context, conn *pgx.Conn, data []model.LoadProfileRecord, dataSetName string) error {
	defer util.TimeTrack(time.Now(), dataSetName, dbName)

	for _, v := range data {
		_, err := conn.Exec(ctx, queryInsertTimeseriesLoadProfileData, v.Ts.Ts, v.LoadProfile, v.Status)
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertPostgresWeatherSingle(ctx context.Context, conn *pgx.Conn, data []model.WeatherRecord, dataSetName string) error {
	defer util.TimeTrack(time.Now(), dataSetName, dbName)

	for _, v := range data {
		num, err := decimal.NewFromString(v.Value)
		if err == nil {
			_, err := conn.Exec(ctx, queryInsertTimeseriesWeatherDataNum, v.Ts.Ts, v.Location, v.Parameter, num, v.Unit)
			if err != nil {
				return err
			}
		} else {
			_, err := conn.Exec(ctx, queryInsertTimeseriesWeatherDataString, v.Ts.Ts, v.Location, v.Parameter, v.Value, v.Unit)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func InsertPostgresLoadProfileBatch(ctx context.Context, conn *pgx.Conn, data *pgx.Batch, dataSetName string) error {
	defer util.TimeTrack(time.Now(), dataSetName, dbName)

	br := conn.SendBatch(ctx, data)
	_, err := br.Exec()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to execute statement in batch queue %v\n", err)
		os.Exit(1)
	}
	return nil
}

func PrepareBatchFromLoadProf(data []model.LoadProfileRecord) *pgx.Batch {
	batch := &pgx.Batch{}
	for _, v := range data {
		batch.Queue(queryInsertTimeseriesLoadProfileData, v.Ts.Ts, v.LoadProfile, v.Status)
	}
	return batch
}
