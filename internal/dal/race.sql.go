// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: race.sql

package dal

import (
	"context"
)

const getRace = `-- name: GetRace :one
SELECT id, season_id, race_name, location, race_date FROM race
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRace(ctx context.Context, id int32) (Race, error) {
	row := q.db.QueryRow(ctx, getRace, id)
	var i Race
	err := row.Scan(
		&i.ID,
		&i.SeasonID,
		&i.RaceName,
		&i.Location,
		&i.RaceDate,
	)
	return i, err
}