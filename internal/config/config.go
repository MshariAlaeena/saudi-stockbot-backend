package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PineconeNamespace    string
	PineconeAPIKey       string
	PineconeHost         string
	PineconeIndex        string
	GroqAPIKey           string
	LLMModel             string
	ArabicLLMModel       string
	MULTIMODAL_LLM_MODEL string
	DBURL                string
	RapidAPIV1Key        string
	RapidAPIV2Key        string
	RapidAPIHost         string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	if os.Getenv("DB_HOST") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_USER") == "" || os.Getenv("DB_PASSWORD") == "" || os.Getenv("DB_NAME") == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}

	sslMode := "disable"
	if os.Getenv("IS_PROD") == "true" {
		sslMode = "require"
	}

	dbURL := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		sslMode,
	)

	cfg := &Config{
		PineconeNamespace:    os.Getenv("PINECONE_NAMESPACE"),
		PineconeAPIKey:       os.Getenv("PINECONE_API_KEY"),
		PineconeIndex:        os.Getenv("PINECONE_INDEX"),
		PineconeHost:         os.Getenv("PINECONE_HOST"),
		GroqAPIKey:           os.Getenv("GROQ_API_KEY"),
		LLMModel:             os.Getenv("LLM_MODEL"),
		ArabicLLMModel:       os.Getenv("ARABIC_LLM_MODEL"),
		MULTIMODAL_LLM_MODEL: os.Getenv("MULTIMODAL_LLM_MODEL"),
		DBURL:                dbURL,
		RapidAPIV1Key:        os.Getenv("RAPID_API_V1_KEY"),
		RapidAPIV2Key:        os.Getenv("RAPID_API_V2_KEY"),
		RapidAPIHost:         os.Getenv("RAPID_API_HOST"),
	}

	if cfg.PineconeAPIKey == "" ||
		cfg.PineconeIndex == "" ||
		cfg.PineconeHost == "" ||
		cfg.GroqAPIKey == "" ||
		cfg.LLMModel == "" ||
		cfg.ArabicLLMModel == "" ||
		cfg.MULTIMODAL_LLM_MODEL == "" ||
		cfg.RapidAPIV1Key == "" ||
		cfg.RapidAPIV2Key == "" ||
		cfg.RapidAPIHost == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}
	return cfg, nil
}
