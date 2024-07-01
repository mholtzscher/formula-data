// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package dal

import (
	"context"
)

type Querier interface {
	CreateSeason(ctx context.Context, arg CreateSeasonParams) (int32, error)
	GetDriver(ctx context.Context, id int32) (Driver, error)
	GetRace(ctx context.Context, id int32) (Race, error)
	GetResult(ctx context.Context, id int32) (Result, error)
	GetSeasonById(ctx context.Context, id int32) (Season, error)
	GetTeam(ctx context.Context, id int32) (Team, error)
	ListTeams(ctx context.Context) ([]Team, error)
}

var _ Querier = (*Queries)(nil)
