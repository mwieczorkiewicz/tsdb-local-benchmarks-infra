package timescale

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func CreateHypertable(ctx context.Context, conn *pgx.Conn, tableName string) error {
	stmt := fmt.Sprintf("SELECT create_hypertable('%s', 'ts')", tableName)
	_, err := conn.Exec(ctx, stmt)
	if err != nil {
		return err
	}
	return nil
}
