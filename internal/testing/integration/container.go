package integration

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	formuladata "github.com/mholtzscher/formula-data"
	"github.com/mholtzscher/formula-data/gen/api/v1/apiv1connect"
	"github.com/mholtzscher/formula-data/internal/dal"
	srvV1 "github.com/mholtzscher/formula-data/internal/service/v1"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type IntegrationTestHelper struct {
	Client           apiv1connect.FormulaDataServiceClient
	Container        *postgres.PostgresContainer
	ConnectionString string
}

func CreateIntegrationTestHelper(t *testing.T) *IntegrationTestHelper {
	ctx := context.Background()
	container, connStr := createPostgresContainer(t, ctx)
	t.Cleanup(func() {
		container.Terminate(ctx)
	})

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("could not connect to database")
	}
	t.Cleanup(func() {
		conn.Close(ctx)
	})

	validator, err := validate.NewInterceptor()
	if err != nil {
		log.Fatal().Err(err).Msg("could not create validation interceptor")
	}

	queries := dal.New(conn)
	fdServer := srvV1.NewFormulaDataServer(queries)

	mux := http.NewServeMux()
	mux.Handle(
		apiv1connect.NewFormulaDataServiceHandler(
			fdServer,
			connect.WithInterceptors(validator),
		),
	)
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	t.Cleanup(server.Close)

	return &IntegrationTestHelper{
		Client:           apiv1connect.NewFormulaDataServiceClient(server.Client(), server.URL),
		Container:        container,
		ConnectionString: connStr,
	}
}

func runMigrations(db *sql.DB) {
	goose.SetBaseFS(formuladata.MigrationsFileSystem)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal().Err(err).Msg("failed to set goose dialect")
	}

	if err := goose.Up(db, "sql/migrations"); err != nil {
		log.Fatal().Err(err).Msg("failed to run migrations on test container")
	}
}

func createPostgresContainer(t *testing.T, ctx context.Context) (*postgres.PostgresContainer, string) {
	t.Helper()
	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:16-alpine"),
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

	runMigrations(conn)

	return pgContainer, connStr
}
