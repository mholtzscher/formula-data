// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package dal

import (
	"context"
)

type Querier interface {
	CreateDriver(ctx context.Context, arg CreateDriverParams) (int32, error)
	CreateRace(ctx context.Context, arg CreateRaceParams) (int32, error)
	CreateResult(ctx context.Context, arg CreateResultParams) (int32, error)
	CreateSeason(ctx context.Context, arg CreateSeasonParams) (int32, error)
	CreateTeam(ctx context.Context, arg CreateTeamParams) (int32, error)
	GetAllSeasons(ctx context.Context) ([]Season, error)
	GetDriverById(ctx context.Context, id int32) (Driver, error)
	GetRaceById(ctx context.Context, id int32) (Race, error)
	GetResultById(ctx context.Context, id int32) (Result, error)
	GetResultsByRaceId(ctx context.Context, raceID int32) ([]Result, error)
	GetSeasonById(ctx context.Context, id int32) (Season, error)
	GetTeamById(ctx context.Context, id int32) (Team, error)
}

var _ Querier = (*Queries)(nil)
