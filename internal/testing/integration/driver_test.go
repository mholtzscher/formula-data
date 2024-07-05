package integration

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/brianvoe/gofakeit/v7"
	apiv1 "github.com/mholtzscher/formula-data/gen/api/v1"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/type/date"
)

func TestCreateDriver(t *testing.T) {
	helper := CreateIntegrationTestHelper(t)
	client := helper.Client

	t.Run("create driver should return season id", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.DriverId)
	})

	t.Run("driver should require first name", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("driver should require last name", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("driver should require place of birth", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("driver should require date of birth", func(t *testing.T) {
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("driver should require date of birth year", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("driver should require date of birth month", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year: int32(d.Year()),
				Day:  int32(d.Day()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("driver should require date of birth day", func(t *testing.T) {
		d := gofakeit.Date()
		result, err := client.CreateDriver(context.Background(), connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
			},
		}))
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeInvalidArgument, connect.CodeOf(err))
		assert.Nil(t, result)
	})

	t.Run("should not allow duplicate driver", func(t *testing.T) {
		d := gofakeit.Date()
		request := connect.NewRequest(&apiv1.CreateDriverRequest{
			FirstName:    gofakeit.FirstName(),
			LastName:     gofakeit.LastName(),
			PlaceOfBirth: gofakeit.City(),
			DateOfBirth: &date.Date{
				Year:  int32(d.Year()),
				Month: int32(d.Month()),
				Day:   int32(d.Day()),
			},
		})

		result, err := client.CreateDriver(context.Background(), request)
		assert.Nil(t, err)
		assert.NotNil(t, result.Msg.DriverId)

		result, err = client.CreateDriver(context.Background(), request)
		assert.NotNil(t, err)
		assert.Equal(t, connect.CodeAlreadyExists, connect.CodeOf(err))
		assert.Nil(t, result)
	})
}
