package config

import (
	"flag"
	"strings"
)

type Config struct {
	Model  string
	Prompt string
}

var Configs = loadConfig()

func loadConfig() *Config {
	var model string
	flag.StringVar(&model, "model", "gpt-3.5-turbo", "GPT model\nOne of: gpt-3.5-turbo, gpt-3.5-turbo-16k, gpt-4, gpt-4-32k")
	flag.Parse()
	prompt := strings.Join(flag.Args(), " ")
	return &Config{
		Model:  model,
		Prompt: prompt,
	}
}
