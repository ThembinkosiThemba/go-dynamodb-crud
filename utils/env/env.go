package env

import "os"

func Getenv(env, defaultValue string) string {

	environment := os.Getenv(env)
	if environment == "" {
		return defaultValue
	}

	return environment
}
