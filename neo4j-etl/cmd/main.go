package main

import (
	"neo4j-etl/internal/db/neo4j"
	"neo4j-etl/internal/handler"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	log := newLog()
	log.Info().Msgf("Running")
	pathReading := os.Getenv("PATH_READING")

	database, err := neo4j.NewNeo4j("bolt://neo4j:7687", "neo4j", "12345678")
	if err != nil {
		log.Fatal().Msgf("error on run server %v", err)
	}

	readInit := handler.NewFileConfig(pathReading, log, database)
	go readInit.ReadCSV()

	time.Sleep(1 * time.Minute)

	select {}
}

func newLog() *zerolog.Logger {
	serviceName := "neo4j-etl"
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
