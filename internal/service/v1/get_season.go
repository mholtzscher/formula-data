package v1

import (
	"context"

	"connectrpc.com/connect"
	apiv1 "github.com/mholtzscher/formula-data/gen/api/v1"
)

func (s *FormulaDataServer) GetSeason(
	ctx context.Context,
	request *connect.Request[apiv1.GetSeasonRequest],
) (*connect.Response[apiv1.GetSeasonResponse], error) {
	season, err := s.DB.GetSeason(ctx, request.Msg.SeasonId)
	if err != nil {
		return nil, err
	}
	return &connect.Response[apiv1.GetSeasonResponse]{
		Msg: &apiv1.GetSeasonResponse{
			Season: &apiv1.Season{
				SeasonId: season.ID,
				Year:     season.SeasonYear,
				Series:   season.Series,
			},
		},
	}, nil
}
