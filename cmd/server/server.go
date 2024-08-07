package main

import (
	"context"
	"database/sql"
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
	"connectrpc.com/vanguard"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	formuladata "github.com/mholtzscher/formula-data"
	"github.com/mholtzscher/formula-data/gen/api/v1/apiv1connect"
	"github.com/mholtzscher/formula-data/internal/dal"
	srvV1 "github.com/mholtzscher/formula-data/internal/service/v1"
	"github.com/peterbourgon/ff/v3"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	fs := flag.NewFlagSet("formula-data", flag.ContinueOnError)
	var (
		listenAddr    = fs.String("listen-addr", "localhost:8080", "listen address")
		logLevel      = fs.String("log-level", "info", "log level")
		dbHost        = fs.String("db-host", "localhost", "database host")
		dbUser        = fs.String("db-user", "postgres", "database user")
		dbPass        = fs.String("db-pass", "postgres", "database password")
		dbName        = fs.String("db-name", "formula-data", "database name")
		dbSslMode     = fs.String("db-sslmode", "prefer", "database sslmode")
		runMigrations = fs.Bool("run-migrations", false, "run database migrations")
	)
	err := ff.Parse(fs, os.Args[1:],
		ff.WithEnvVarPrefix("FORMULA_DATA"),
		ff.WithConfigFile(".env"),
		ff.WithAllowMissingConfigFile(true),
		ff.WithConfigFileParser(ff.EnvParser),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("error parsing flags")
	}

	setupLogging(*logLevel)

	ctx := context.Background()

	log.Info().Str("host", *dbHost).Str("user", *dbUser).Msg("connecting to database")
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", *dbHost, *dbUser, *dbPass, *dbName, *dbSslMode)

	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl != "" {
		connString = dbUrl
		log.Info().Msg("using DATABASE_URL for postgres connection")
	}

	if *runMigrations {
		log.Info().Msg("running migrations")
		err = runGooseMigrations(connString)
		if err != nil {
			log.Fatal().Err(err).Msg("could not run migrations")
		}
		return
	}

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

	service := vanguard.NewService(apiv1connect.NewFormulaDataServiceHandler(
		fdServer,
		connect.WithInterceptors(validator),
	),
	)

	handler, err := vanguard.NewTranscoder([]*vanguard.Service{service})
	if err != nil {
		log.Fatal().Err(err).Msg("could not create transcoder")
	}

	srv := &http.Server{
		Addr: *listenAddr,
		Handler: h2c.NewHandler(
			handler,
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
		log.Info().Msg("starting server")
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

func runGooseMigrations(connString string) error {
	goose.SetBaseFS(formuladata.MigrationsFileSystem)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	db, err := sql.Open("pgx", connString)
	if err != nil {
		return err
	}
	defer db.Close()
	if err := goose.Up(db, "sql/migrations"); err != nil {
		return err
	}
	return err
}
