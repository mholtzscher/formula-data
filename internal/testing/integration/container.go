package integration

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	formuladata "github.com/mholtzscher/formula-data"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	ConnectionString string
}

func RunMigrations(db *sql.DB) {
	goose.SetBaseFS(formuladata.MigrationsFileSystem)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal().Err(err).Msg("failed to set goose dialect")
	}

	if err := goose.Up(db, "sql/migrations"); err != nil {
		log.Fatal().Err(err).Msg("failed to run migrations on test container")
	}
}

func CreatePostgresContainer(ctx context.Context) *PostgresContainer {
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:15.3-alpine"),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start postgres test container")
	}

	connStr, err := pgContainer.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get connection string from test container")
	}

	conn, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open sql conn to test container")
	}
	defer conn.Close()

	RunMigrations(conn)

	return &PostgresContainer{
		PostgresContainer: pgContainer,
		ConnectionString:  connStr,
	}
}
