package config

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"terminal-ai/color"
)

type Config struct {
	Model  string
	Prompt string
}

var Configs = loadConfig()

func loadConfig() *Config {
	var model string
	modelDesc := fmt.Sprintf("%s: %s\n", color.Format(color.CYAN, "Full list of models"), color.Format(color.BLUE, "https://platform.openai.com/docs/models"))
	flag.StringVar(&model, "model", "gpt-3.5-turbo", modelDesc)

	flag.Parse()

	prompt := strings.Join(flag.Args(), " ")
	if prompt == "" {
		log.Fatal("Please provide a prompt")
		return nil
	}
	return &Config{
		Model:  model,
		Prompt: prompt,
	}
}
