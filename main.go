package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"ratelimit/config"
	"ratelimit/server"
)

func init() {
	config.Init()
	server.InitRedis()
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	server.Run()
}
