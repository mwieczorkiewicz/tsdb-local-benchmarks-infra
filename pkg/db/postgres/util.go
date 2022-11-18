package postgres

import (
	"context"
	"io/ioutil"

	pgx "github.com/jackc/pgx/v5"
)

func SetupTablesPostgres(ctx context.Context, conn *pgx.Conn, path string) error {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	sql := string(c)
	_, err = conn.Exec(ctx, sql)
	if err != nil {
		return err
	}
	return nil
}
