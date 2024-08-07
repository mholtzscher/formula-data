package v1

import (
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgtype"
	apiv1 "github.com/mholtzscher/formula-data/gen/api/v1"
	"github.com/mholtzscher/formula-data/internal/dal"
	"github.com/mholtzscher/formula-data/internal/testing/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/genproto/googleapis/type/date"
)

func TestCreateDriver(t *testing.T) {
	mockDB := mocks.NewMockQuerier(t)
	service := NewFormulaDataServer(mockDB)

	t.Run("should create driver", func(t *testing.T) {
		mockDB.On("CreateDriver", mock.Anything, dal.CreateDriverParams{
			FirstName:    "Max",
			LastName:     "Verstappen",
			PlaceOfBirth: "Hasselt, Belgium",
			DateOfBirth: pgtype.Date{
				Time:  time.Date(1997, 9, 30, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
		}).Return(int32(1), nil).Once()

		request := &connect.Request[apiv1.CreateDriverRequest]{
			Msg: &apiv1.CreateDriverRequest{
				FirstName:    "Max",
				LastName:     "Verstappen",
				PlaceOfBirth: "Hasselt, Belgium",
				DateOfBirth: &date.Date{
					Year:  1997,
					Month: 9,
					Day:   30,
				},
			},
		}

		result, err := service.CreateDriver(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, int32(1), result.Msg.DriverId)
	})

	t.Run("should return error when create driver fails", func(t *testing.T) {
		mockDB.On("CreateDriver", mock.Anything, mock.Anything).Return(int32(0), assert.AnError).Once()

		request := &connect.Request[apiv1.CreateDriverRequest]{
			Msg: &apiv1.CreateDriverRequest{
				FirstName:    "Max",
				LastName:     "Verstappen",
				PlaceOfBirth: "Hasselt, Belgium",
				DateOfBirth: &date.Date{
					Year:  1997,
					Month: 9,
					Day:   30,
				},
			},
		}
		result, err := service.CreateDriver(context.Background(), request)

		mockDB.AssertExpectations(t)
		assert.NotNil(t, err)
		assert.Nil(t, result)
	})
}
