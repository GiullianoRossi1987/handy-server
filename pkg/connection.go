package pkg

import (
	"context"
	"fmt"
	types "types/config"

	"github.com/jackc/pgx/v5"
)

func generateConnection(config types.PsConfig) *pgx.Conn {
	pSqlInfo := fmt.Sprintf("host=%s;username=%s;password=%s;db=%s", config.Host, config.Username, config.Password, config.Db)
	psql, err := pgx.Connect(context.Background(), pSqlInfo)
	if err != nil {
		panic(err)
	}
	return psql
}
