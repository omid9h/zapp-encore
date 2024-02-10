package errorhandler

import (
	"errors"
	"fmt"

	"encore.dev/beta/errs"
	"github.com/jackc/pgx/v5/pgconn"
)

func HandleDBError(err error) error {
	if err == nil {
		return nil
	}
	var (
		newerr errs.Error
		pgErr  *pgconn.PgError
	)
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			newerr.Code = errs.AlreadyExists
			newerr.Message = fmt.Sprintf("duplicate key value violates unique constraint %s", pgErr.ConstraintName)
		case "23503":
			newerr.Code = errs.FailedPrecondition
			newerr.Message = fmt.Sprintf("foreign key constraint %s violated", pgErr.ConstraintName)
		case "23502":
			newerr.Code = errs.FailedPrecondition
			newerr.Message = fmt.Sprintf("null value in column %s violates not-null constraint", pgErr.ConstraintName)
		default:
			newerr.Code = errs.Aborted
			newerr.Message = pgErr.Message
		}
	}
	return &newerr
}
