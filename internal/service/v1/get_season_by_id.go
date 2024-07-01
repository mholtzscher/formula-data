package v1

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/mholtzscher/formula-data/gen/api/v1"
)

func (s *FormulaDataServer) GetSeasonById(
	ctx context.Context,
	request *connect.Request[apiv1.GetSeasonByIdRequest],
) (*connect.Response[apiv1.GetSeasonByIdResponse], error) {
	season, err := s.DB.GetSeasonById(ctx, request.Msg.SeasonId)
	if err != nil {
		return nil, err
	}
	return &connect.Response[apiv1.GetSeasonByIdResponse]{
		Msg: &apiv1.GetSeasonByIdResponse{
			Season: &apiv1.Season{
				SeasonId: season.ID,
				Year:     season.SeasonYear,
				Series:   season.Series,
			},
		},
	}, nil
}
