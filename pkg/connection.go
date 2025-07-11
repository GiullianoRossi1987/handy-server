package pkg

import (
	"context"
	"fmt"
	types "types/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DEPRECATE
func GenerateConnection(config types.PsConfig) (*pgx.Conn, error) {
	pSqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:6611/%s", config.Username, config.Password, config.Host, config.Db)
	psql, err := pgx.Connect(context.Background(), pSqlInfo)
	if err != nil {
		return nil, err
	}
	return psql, nil
}

func GeneratePool(config types.PsConfig) (*pgxpool.Pool, error) {
	pSqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:6611/%s", config.Username, config.Password, config.Host, config.Db)
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
