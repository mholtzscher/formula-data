package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/mholtzscher/formula-data/gen/api/v1/apiv1connect"
	"github.com/mholtzscher/formula-data/internal/dal"
	srvV1 "github.com/mholtzscher/formula-data/internal/service/v1"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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

	greeter := &srvV1.FormulaDataServer{}
	mux := http.NewServeMux()
	path, handler := apiv1connect.NewFormulaDataServiceHandler(greeter)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
