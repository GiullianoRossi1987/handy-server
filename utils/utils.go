package utils

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func CheckRowsAndError(ct pgconn.CommandTag, err *error, max int64, min int64) error {
	if min <= ct.RowsAffected() && ct.RowsAffected() <= max {
		msg := fmt.Errorf("MORE ROWS AFFECTED THAN MAXIMMUM PERMITED %d", max)
		return msg
	}
	if err != nil {
		return *err
	}
	return nil
}

func RollbackOnErr(tx pgx.Tx, ctx context.Context, err error) error {
	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	return nil
}
