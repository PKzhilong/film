package config

import "os"

var (
	MYSQL_HOST = getEnv("MYSQL_HOST", "")
	MYSQL_USER = getEnv("MYSQL_USER", "root")
	MYSQL_PASSWORD = getEnv("MYSQL_PASSWORD", "admin")
	MYSQL_DATABASE = getEnv("MYSQL_DATABASE", "")
)

func getEnv(key string, defaultValue interface{}) interface{} {
	result := os.Getenv(key)

	if len(result) < 1 {
		return defaultValue
	}
	return result
}