package services

import (
	db "pkg"
	config "types/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func GetConnByEnv() (*pgxpool.Pool, error) {
	cn := config.PsConfig{}
	cn.FromEnv()
	return db.GeneratePool(cn)
}
