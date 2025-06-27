package utils

import (
	"os"
	types "types/config"
)

func GenerateDatabaseConfig() (*types.PsConfig, error) {
	return &types.PsConfig{
		Host:     os.Getenv("HOSTNAME"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Db:       os.Getenv("DATABASE"),
	}, nil
}
