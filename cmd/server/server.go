package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/mholtzscher/formula-data/gen/api/v1/apiv1connect"
	"github.com/mholtzscher/formula-data/internal/dal"
	srvV1 "github.com/mholtzscher/formula-data/internal/service/v1"
	"github.com/peterbourgon/ff/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	fs := flag.NewFlagSet("formula-data", flag.ContinueOnError)
	var (
		listenAddr = fs.String("listen-addr", "localhost:8080", "listen address")
		logLevel   = fs.String("log-level", "info", "log level")
	)
	err := ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("FORMULA_DATA"),
		ff.WithConfigFile(".env"),
		ff.WithConfigFileParser(ff.EnvParser),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("error parsing flags")
	}

	setupLogging(*logLevel)

	log.Info().Msg("starting server")

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
		*listenAddr,
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}

func setupLogging(logLevel string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	switch strings.ToLower(logLevel) {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	}
}
