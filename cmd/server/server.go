package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/mholtzscher/formula-data/gen/api/v1/apiv1connect"
	"github.com/mholtzscher/formula-data/internal/dal"
	srvV1 "github.com/mholtzscher/formula-data/internal/service/v1"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var debug = flag.Bool("debug", false, "sets log level to debug")

func main() {
	flag.Parse()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	log.Info().Msg("Starting server...")

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "host=localhost user=postgres password=postgres dbname=formula-data")
	if err != nil {
		log.Fatal().Err(err).Msg("could not connect to database")
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
}
