package config

import (
	"log"
	"os"

	types "types/config"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
	return os.Getenv(key)
}

func GetConfigByEnv() types.PsConfig {
	hostname := GetEnv("HOSTNAME")
	username := GetEnv("USERNAME")
	password := GetEnv("PASSWORD")
	database := GetEnv("DATABASE")
	return types.PsConfig{
		Host:     hostname,
		Username: username,
		Password: password,
		Db:       database,
	}
}
