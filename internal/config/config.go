package config

import (
	"fmt"
	"os"
)

type Config struct {
	OpenAIAPIKey string
	Model        string
}

func Load() (*Config, error) {
	apikey := os.Getenv("OPENAI_API_KEY")
	if apikey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY is required")
	}
	model := os.Getenv("OPENAI_MODEL")
	if model == "" {
		model = "gpt-4.1-mini"
	}
	return &Config{
		OpenAIAPIKey: apikey,
		Model:        model,
	}, nil

}
