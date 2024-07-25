package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"connectrpc.com/vanguard"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mholtzscher/formula-data/gen/api/v1/apiv1connect"
	"github.com/mholtzscher/formula-data/internal/dal"
	srvV1 "github.com/mholtzscher/formula-data/internal/service/v1"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Getenv); err != nil {
		log.Fatal().Err(err).Msg("error running server")
	}
}

func run(ctx context.Context, getEnv func(string) string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	dbFile := getEnv("DB_FILE")
	host := getEnv("HOST")
	port := getEnv("PORT")
	logLevel := getEnv("LOG_LEVEL")

	setupLogging(logLevel)

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return err
	}
	defer db.Close()
	log.Info().Str("db-file", dbFile).Msg("connected to sqlite database")

	queries := dal.New(db)
	fdServer := srvV1.NewFormulaDataServer(queries)

	validator, err := validate.NewInterceptor()
	if err != nil {
		return err
	}

	service := vanguard.NewService(apiv1connect.NewFormulaDataServiceHandler(
		fdServer,
		connect.WithInterceptors(validator),
	),
	)

	handler, err := vanguard.NewTranscoder([]*vanguard.Service{service})
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr: net.JoinHostPort(host, port),
		Handler: h2c.NewHandler(
			handler,
			&http2.Server{},
		),
		ReadHeaderTimeout: time.Second,
		ReadTimeout:       5 * time.Minute,
		WriteTimeout:      5 * time.Minute,
		MaxHeaderBytes:    8 * 1024, // 8KiB
	}

	go func() {
		log.Info().Msg("starting server")
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		// make a new context for the Shutdown (thanks Alessandro Rosetti)
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Error().Err(err).Msg("error shutting down http server")
		}
	}()
	wg.Wait()
	return nil
}

// func main() {
// 	fs := flag.NewFlagSet("formula-data", flag.ContinueOnError)
// 	var (
// 		listenAddr = fs.String("listen-addr", "localhost:8080", "listen address")
// 		logLevel   = fs.String("log-level", "info", "log level")
// 		dbFile     = fs.String("db-file", "f1db.db", "database file")
// 	)
// 	err := ff.Parse(fs, os.Args[1:],
// 		ff.WithEnvVarPrefix("FORMULA_DATA"),
// 		ff.WithConfigFile(".env"),
// 		ff.WithAllowMissingConfigFile(true),
// 		ff.WithConfigFileParser(ff.EnvParser),
// 	)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("error parsing flags")
// 	}
//
// 	setupLogging(*logLevel)
//
// 	db, err := sql.Open("sqlite3", *dbFile)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("could not open database")
// 	}
// 	defer db.Close()
// 	log.Info().Str("db-file", *dbFile).Msg("connected to sqlite database")
//
// 	queries := dal.New(db)
// 	fdServer := srvV1.NewFormulaDataServer(queries)
//
// 	validator, err := validate.NewInterceptor()
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("could not create validation interceptor")
// 	}
//
// 	service := vanguard.NewService(apiv1connect.NewFormulaDataServiceHandler(
// 		fdServer,
// 		connect.WithInterceptors(validator),
// 	),
// 	)
//
// 	handler, err := vanguard.NewTranscoder([]*vanguard.Service{service})
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("could not create transcoder")
// 	}
//
// 	srv := &http.Server{
// 		Addr: *listenAddr,
// 		Handler: h2c.NewHandler(
// 			handler,
// 			&http2.Server{},
// 		),
// 		ReadHeaderTimeout: time.Second,
// 		ReadTimeout:       5 * time.Minute,
// 		WriteTimeout:      5 * time.Minute,
// 		MaxHeaderBytes:    8 * 1024, // 8KiB
// 	}
//
// 	signals := make(chan os.Signal, 1)
// 	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
// 	go func() {
// 		log.Info().Msg("starting server")
// 		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
// 			log.Fatal().Msgf("HTTP listen and serve: %v", err)
// 		}
// 	}()
//
// 	<-signals
// 	log.Info().Msg("shutting down server")
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()
// 	if err := srv.Shutdown(ctx); err != nil {
// 		log.Fatal().Msgf("HTTP shutdown: %v", err)
// 	}
// }

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
