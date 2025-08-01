package main

import (
	"patient-chatbot/internal/config"
	"patient-chatbot/internal/utils"

	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Error().Msg("error loading config: " + err.Error())
	}

	utils.Init()
	server := NewServer(cfg)
	log.Info().Msg("Starting server...")

	if err := server.Run(); err != nil {
		log.Error().Msg("error running server: " + err.Error())
	}
}
