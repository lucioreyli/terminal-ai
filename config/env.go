package config

import (
	"os"
)

type Env struct {
	MODEL          string
	OPENAI_API_KEY string
}

func loadEnv() *Env {
	return &Env{
		MODEL:          getEnv("MODEL", "gpt-3.5-turbo"),
		OPENAI_API_KEY: getEnv("GPT_KEY", ""),
	}
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

var Envs = loadEnv()
