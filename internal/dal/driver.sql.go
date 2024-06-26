// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: driver.sql

package dal

import (
	"context"
)

const getDriver = `-- name: GetDriver :one
SELECT id, first_name, last_name, place_of_birth, date_of_birth FROM driver
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetDriver(ctx context.Context, id int32) (Driver, error) {
	row := q.db.QueryRow(ctx, getDriver, id)
	var i Driver
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.PlaceOfBirth,
		&i.DateOfBirth,
	)
	return i, err
}
