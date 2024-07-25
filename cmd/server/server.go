package main

import (
	"context"
	"database/sql"
	"errors"
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
	"github.com/sethvargo/go-envconfig"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Config struct {
	Host     string `env:"HOST, default=localhost"`
	Port     string `env:"PORT, default=8080"`
	LogLevel string `env:"LOG_LEVEL, default=info"`
	DbFile   string `env:"DB_FILE, default=f1db.db"`
}

func main() {
	ctx := context.Background()
	if err := run(ctx, envconfig.OsLookuper()); err != nil {
		log.Fatal().Err(err).Msg("error running server")
	}
}

func run(ctx context.Context, envLookup envconfig.Lookuper) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	var cfg Config
	err := envconfig.ProcessWith(ctx, &envconfig.Config{
		Target:   &cfg,
		Lookuper: envLookup,
	})
	if err != nil {
		return err
	}
	log.Info().Interface("config", cfg).Msg("loaded configuration")

	setupLogging(cfg.LogLevel)

	db, err := sql.Open("sqlite3", cfg.DbFile)
	if err != nil {
		return err
	}
	defer db.Close()
	log.Info().Msg("connected to sqlite database")

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
		Addr: net.JoinHostPort(cfg.Host, cfg.Port),
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
			log.Error().Err(err).Msg("error listening and serving")
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
	log.Info().Msg("shutdown complete")
	return nil
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
