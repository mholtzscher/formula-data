// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package dal

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Driver struct {
	ID           int32
	FirstName    string
	LastName     string
	PlaceOfBirth string
	DateOfBirth  pgtype.Date
}

type Season struct {
	ID         int32
	SeasonYear int32
	Series     string
}
