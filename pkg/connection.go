package pkg

import (
	"context"
	"fmt"
	types "types/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GeneratePool(config types.PsConfig) (*pgxpool.Pool, error) {
	pSqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", config.Username, config.Password, config.Host, config.Port, config.Db)
	conf, err := pgxpool.ParseConfig(pSqlInfo)
	if err != nil {
		return nil, err
	}
	psql, err := pgxpool.NewWithConfig(context.Background(), conf)
	if err != nil {
		return nil, err
	}
	return psql, nil
}
