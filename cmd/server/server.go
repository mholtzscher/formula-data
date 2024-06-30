package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
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
		dbHost     = fs.String("db-host", "localhost", "database host")
		dbUser     = fs.String("db-user", "postgres", "database user")
		dbPass     = fs.String("db-pass", "postgres", "database password")
		dbName     = fs.String("db-name", "formula-data", "database name")
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

	ctx := context.Background()

	log.Info().Str("host", *dbHost).Str("user", *dbUser).Msg("connecting to database")
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", *dbHost, *dbUser, *dbPass, *dbName)
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatal().Err(err).Msg("could not connect to database")
	}
	defer conn.Close(ctx)
	log.Info().Msg("connected to postgres database")

	queries := dal.New(conn)
	fdServer := srvV1.NewFormulaDataServer(queries)

	validator, err := validate.NewInterceptor()
	if err != nil {
		log.Fatal().Err(err).Msg("could not create validation interceptor")
	}

	log.Info().Msg("starting server")
	mux := http.NewServeMux()
	mux.Handle(apiv1connect.NewFormulaDataServiceHandler(
		fdServer,
		connect.WithInterceptors(validator),
	),
	)

	srv := &http.Server{
		Addr: *listenAddr,
		Handler: h2c.NewHandler(
			mux,
			&http2.Server{},
		),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Msgf("HTTP listen and serve: %v", err)
		}
	}()

	<-signals
	log.Info().Msg("shutting down server")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msgf("HTTP shutdown: %v", err)
	}
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
