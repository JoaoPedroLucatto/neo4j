package main

import (
	"context"
	"neo4j-api/internal/db/neo4j"
	"neo4j-api/internal/handler"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	log := newLog()
	log.Info().Msgf("Running")

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPswd := os.Getenv("DB_PASSWORD")

	database, err := neo4j.NewNeo4j(dbHost, dbUser, dbPswd)
	if err != nil {
		log.Fatal().Msgf("error on run server %v", err)

		return
	}

	s := handler.NewServer(context.Background(), log, database)
	if err := s.Server().Run(os.Getenv("HOST")); err != nil {
		log.Fatal().Msgf("error on run server %v", err)
	}
}

func newLog() *zerolog.Logger {
	serviceName := "neo4j-api"
	log := zerolog.New(os.Stdout).With().
		Timestamp().Str("service", serviceName).Logger()
	logLevel := "debug"

	loggerLevel, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		loggerLevel = zerolog.DebugLevel

		log.Warn().Msgf("logger level invalid '%s'. Using logger level default 'info'.", logLevel)
	}

	zerolog.SetGlobalLevel(loggerLevel)

	return &log
}
