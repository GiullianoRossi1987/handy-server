package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}
	return os.Getenv(key)
}

type PsConfig struct {
	Host     string
	Username string
	Password string
	Db       string
}

func GetConfigByEnv() PsConfig {
	hostname := GetEnv("HOSTNAME")
	username := GetEnv("USERNAME")
	password := GetEnv("PASSWORD")
	database := GetEnv("DB")
	return PsConfig{
		Host:     hostname,
		Username: username,
		Password: password,
		Db:       database,
	}
}
