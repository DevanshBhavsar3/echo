package config

import "os"

type Config struct {
	REDIS_URL string
	REGION    string
	WORKER_ID string
}

func LoadEnv() Config {
	return Config{
		REDIS_URL: GetEnv("REDIS_URL", "localhost:6379"),
		REGION:    GetEnv("REGION", "UNKNOWN"),
		WORKER_ID: GetEnv("WORKER_ID", "1"),
	}
}

func GetEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		return defaultValue
	}

	return value
}
