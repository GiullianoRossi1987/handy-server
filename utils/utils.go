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

// fucking hate the fact there there isn't FP in this language, why google why
func MapCar[T any, U any](ia []T, of func(item T) U) []U {
	mapped := make([]U, len(ia))
	for k, i := range ia {
		mapped[k] = of(i)
	}
	return mapped
}

func FilterCar[T any](pool []T, check func(i T) bool) []T {
	filtered := []T{}
	for _, i := range pool {
		if check(i) {
			filtered = append(filtered, i)
		}
	}
	return filtered
}
