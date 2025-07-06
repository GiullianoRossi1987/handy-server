package utils

import (
	"os"
	types "types/config"
)

func GenerateDatabaseConfig() types.PsConfig {
	return types.PsConfig{
		Host:     os.Getenv("HOSTNAME"),
		Username: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		Db:       os.Getenv("DATABASE"),
	}
}

// Why GO doesnt ship without a '??' coalesce function? I gotta do this ¬~¬
func Coalesce[T any](value *T, default_val T) T {
	if value == nil {
		return default_val
	}
	return *value
}
