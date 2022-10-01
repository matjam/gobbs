package main

import (
	"fmt"
	"net"
	"os"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/pborman/ansi"
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

	telnetPort := config.Int("telnet.port", 2125)

	log.Info().Msgf("telnet service starting on port %v ...", telnetPort)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", telnetPort))
	if err != nil {
		log.Fatal().Msgf("couldn't listen to telnet port %v", err.Error())
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal().Msgf("couldn't accept connection from %v: %v", conn.RemoteAddr(), err.Error())
		}
		go handleTelnetConnection(conn)
	}
}

func handleTelnetConnection(conn net.Conn) {
	log.Info().Msgf("[%v] accepted connection via telnet", conn.RemoteAddr())
	c := ansi.NewWriter(conn)
	c.Colorize()
	c.ForceReset()
	c.Blue().Bold().WriteString(config.String("telnet.banner", "GOBBS"))
	conn.Close()
}
