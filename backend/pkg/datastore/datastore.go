package datastore

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

var pg *pgxpool.Pool

func Init(ctx context.Context) error {
	dbpool, err := pgxpool.New(ctx, "")
	if err != nil {
		return fmt.Errorf("pgxpool.New: %w", err)
	}

	if err := dbpool.Ping(ctx); err != nil {
		return fmt.Errorf("dbpool.Ping: %w", err)
	}

	pg = dbpool
	return nil
}

func Shutdown() {
	pg.Close()
}
