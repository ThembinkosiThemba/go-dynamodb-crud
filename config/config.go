package config

import (
	"strconv"

	"github.com/ThembinkosiThemba/dynamodb-go-crud/utils/env"
)

type Config struct {
	Port, Timeout        int
	Dialect, DatabaseURI string
}

func GetConfig() Config {
	return Config{
		Port:        parseEnvToInteger("PORT", "8080"),
		Timeout:     parseEnvToInteger("TIMEOUT", "20"),
		Dialect:     env.Getenv("DIALECT", "sqlite3"),
		DatabaseURI: env.Getenv("DATABASE_URI", ":memory"),
	}
}

func parseEnvToInteger(envName, defaultValues string) int {
	num, err := strconv.Atoi(env.Getenv(envName, defaultValues))
	if err != nil {
		return 0
	}

	return num
}
