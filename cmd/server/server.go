package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/mholtzscher/formula-data/internal/dal"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=localhost user=postgres password=postgres dbname=formula-data")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := dal.New(conn)
	_ = queries

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
