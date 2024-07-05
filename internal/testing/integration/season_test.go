package integration

import (
	"context"
	"testing"

	apiv1 "github.com/mholtzscher/formula-data/gen/api/v1"

	"connectrpc.com/connect"
	"github.com/stretchr/testify/assert"
)

func TestFormulaDataServer_Season(t *testing.T) {
	client := CreateTestServerAndClient(t)

	t.Run("create season should return season id", func(t *testing.T) {
		result, err := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
			Year:   2024,
			Series: "Formula Test",
		}))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.SeasonId)
	})

	t.Run("season should require year", func(t *testing.T) {
		result, err := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
			Series: "Formula Test",
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("season should be after 1900", func(t *testing.T) {
		result, err := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
			Year:   1900,
			Series: "Formula Test",
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("season should be before 2100", func(t *testing.T) {
		result, err := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
			Year:   2101,
			Series: "Formula Test",
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("season should require series name", func(t *testing.T) {
		result, err := client.CreateSeason(context.Background(), connect.NewRequest(&apiv1.CreateSeasonRequest{
			Year: 2024,
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("should not allow duplicate season", func(t *testing.T) {
		request := connect.NewRequest(&apiv1.CreateSeasonRequest{
			Series: "Formula Test 123",
			Year:   2024,
		})

		result, err := client.CreateSeason(context.Background(), request)
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.SeasonId)

		result, err = client.CreateSeason(context.Background(), request)
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeAlreadyExists, connect.CodeOf(err))
		assert.Nil(t, result)
	})
}
