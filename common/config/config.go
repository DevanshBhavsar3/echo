package config

import (
	"log"

	"github.com/joho/godotenv"
)

var cfg map[string]string

func init() {
	var err error
	cfg, err = godotenv.Read()
	if err != nil {
		log.Fatalf("error loading .env:\n%v", err)
	}
}

func Get(key string) string {
	return cfg[key]
}
