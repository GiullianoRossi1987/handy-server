package utils

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

func CheckRowsAndError(ct pgconn.CommandTag, err *error, max int64) {
	if ct.RowsAffected() > max {
		msg := fmt.Sprintf("MORE ROWS AFFECTED THAN MAXIMMUM PERMITED %d", max)
		panic(msg)
	}
	if err != nil {
		panic(err)
	}
}
