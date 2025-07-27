package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GroqAPIKey           string
	LLMModel             string
	ArabicLLMModel       string
	MULTIMODAL_LLM_MODEL string
	RapidAPIV1Key        string
	RapidAPIV2Key        string
	RapidAPIHost         string
	FrontendURL          string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	cfg := &Config{
		GroqAPIKey:           os.Getenv("GROQ_API_KEY"),
		LLMModel:             os.Getenv("LLM_MODEL"),
		ArabicLLMModel:       os.Getenv("ARABIC_LLM_MODEL"),
		MULTIMODAL_LLM_MODEL: os.Getenv("MULTIMODAL_LLM_MODEL"),
		RapidAPIV1Key:        os.Getenv("RAPID_API_V1_KEY"),
		RapidAPIV2Key:        os.Getenv("RAPID_API_V2_KEY"),
		RapidAPIHost:         os.Getenv("RAPID_API_HOST"),
		FrontendURL:          os.Getenv("FRONTEND_URL"),
	}

	if cfg.GroqAPIKey == "" ||
		cfg.LLMModel == "" ||
		cfg.ArabicLLMModel == "" ||
		cfg.MULTIMODAL_LLM_MODEL == "" ||
		cfg.RapidAPIV1Key == "" ||
		cfg.RapidAPIV2Key == "" ||
		cfg.RapidAPIHost == "" ||
		cfg.FrontendURL == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}
	return cfg, nil
}
