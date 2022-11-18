package postgres

import (
	"context"
	"os"

	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
	"github.com/jackc/pgx/v5"
)

func NewPostgresConnection() (*pgx.Conn, error) {
	// sample: "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	pgxdecimal.Register(conn.TypeMap())
	return conn, err
}
