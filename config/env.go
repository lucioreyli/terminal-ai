package config

import (
	"os"
)

type Env struct {
	OPENAI_API_KEY string
	GPT_URL        string
}

func loadEnv() *Env {
	return &Env{
		OPENAI_API_KEY: getEnv("GPT_KEY", ""),
		GPT_URL:        getEnv("GPT_URL", "https://api.openai.com/v1/chat/completions"),
	}
}

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

var Envs = loadEnv()
