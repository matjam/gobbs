package main

import (
	"os"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
	"github.com/matjam/gobbs/internal/servers"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("starting GOBBS server")

	config.WithOptions(config.ParseEnv)

	// add driver for support yaml content
	config.AddDriver(toml.Driver)

	err := config.LoadFiles("gobbs.toml")
	if err != nil {
		log.Fatal().Msgf("couldn't open configuration gobbs.toml: %v", err.Error())
	}

	servers.Start()
}
