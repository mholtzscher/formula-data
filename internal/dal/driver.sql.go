// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: driver.sql

package dal

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createDriver = `-- name: CreateDriver :one
INSERT INTO driver
(first_name, last_name, place_of_birth, date_of_birth)
VALUES (
$1, $2, $3, $4
)
RETURNING id
`

type CreateDriverParams struct {
	FirstName    string
	LastName     string
	PlaceOfBirth string
	DateOfBirth  pgtype.Date
}

func (q *Queries) CreateDriver(ctx context.Context, arg CreateDriverParams) (int32, error) {
	row := q.db.QueryRow(ctx, createDriver,
		arg.FirstName,
		arg.LastName,
		arg.PlaceOfBirth,
		arg.DateOfBirth,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getDriverById = `-- name: GetDriverById :one
SELECT id, first_name, last_name, place_of_birth, date_of_birth FROM driver
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetDriverById(ctx context.Context, id int32) (Driver, error) {
	row := q.db.QueryRow(ctx, getDriverById, id)
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
