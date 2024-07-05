package v1

import (
	"errors"

	"connectrpc.com/connect"
	"github.com/jackc/pgx/v5/pgconn"
)

func mapPgErrorsToReturnCodes(err error) *connect.Error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return connect.NewError(connect.CodeAlreadyExists, pgErr)
		default:
			return connect.NewError(connect.CodeUnknown, pgErr)
		}
	}
	return connect.NewError(connect.CodeUnknown, pgErr)
}
