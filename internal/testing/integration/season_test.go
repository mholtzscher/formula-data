package integration

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jackc/pgx/v5"
	apiv1 "github.com/mholtzscher/formula-data/gen/api/v1"
	"github.com/mholtzscher/formula-data/gen/api/v1/apiv1connect"
	"github.com/rs/zerolog/log"

	"github.com/mholtzscher/formula-data/internal/dal"
	srvV1 "github.com/mholtzscher/formula-data/internal/service/v1"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormulaDataServer_Season(t *testing.T) {
	ctx := context.Background()
	container := CreatePostgresContainer(ctx)
	conn, err := pgx.Connect(ctx, container.ConnectionString)
	if err != nil {
		log.Fatal().Err(err).Msg("could not connect to database")
	}
	defer conn.Close(ctx)

	queries := dal.New(conn)
	fdServer := srvV1.NewFormulaDataServer(queries)

	mux := http.NewServeMux()
	mux.Handle(apiv1connect.NewFormulaDataServiceHandler(fdServer))
	server := httptest.NewUnstartedServer(mux)
	server.EnableHTTP2 = true
	server.StartTLS()
	defer server.Close()

	connectClient := apiv1connect.NewFormulaDataServiceClient(server.Client(), server.URL)
	grpcClient := apiv1connect.NewFormulaDataServiceClient(server.Client(), server.URL, connect.WithGRPC())
	clients := []apiv1connect.FormulaDataServiceClient{connectClient, grpcClient}

	t.Run("create season", func(t *testing.T) {
		for _, client := range clients {
			result, err := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
				Year:   2024,
				Series: "Formula Test",
			}))
			require.Nil(t, err)
			assert.NotNil(t, result.Msg.SeasonId)
		}
	})
}
