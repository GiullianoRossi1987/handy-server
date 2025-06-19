package pkg

import (
	"context"
	"fmt"
	types "types/config"

	"github.com/jackc/pgx/v5"
)

func GenerateConnection(config types.PsConfig) *pgx.Conn {
	pSqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:6611/%s", config.Username, config.Password, config.Host, config.Db)
	psql, err := pgx.Connect(context.Background(), pSqlInfo)
	if err != nil {
		panic(err)
	}
	return psql
}
