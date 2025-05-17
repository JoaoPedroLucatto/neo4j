package handler

import (
	"encoding/csv"
	"io"
	"neo4j-etl/internal/db"
	"neo4j-etl/internal/usecase"
	"neo4j-etl/internal/util"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Usecase struct {
	Usecase *usecase.Usecase
	Path    string
	Logger  *zerolog.Logger
}

func NewFileConfig(path string, logger *zerolog.Logger, repository db.Repository) *Usecase {
	logger.Info().Str("Read", "ReadInit").Msgf("Initializing read CSV %s", path)

	return &Usecase{
		Usecase: usecase.NewUsecaseService(logger, repository),
		Path:    path,
		Logger:  logger,
	}
}

func (usecase *Usecase) ReadCSV() error {
	log := usecase.Logger.With().Str("handler", "ReadCSV").Logger()

	f, err := os.Open(usecase.Path)
	if err != nil {
		log.Err(err).Msgf("erro file open %v", err)
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	err = processCSV(reader, usecase)
	if err != nil {
		log.Err(err).Msgf("erro process CSV: %v", err)
		return err
	}

	return nil
}

func processCSV(reader *csv.Reader, uc *Usecase) error {
	header, err := reader.Read()
	if err != nil {
		uc.Logger.Err(err).Msgf("erro ao ler o cabeçalho do arquivo %s: %v", uc.Path, err)
		return err
	}
	uc.Logger.Info().Msgf("Cabeçalho lido: %v", header)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			uc.Logger.Err(err).Msgf("erro process read in file %s: %v", uc.Path, err)
			return err
		}

		if err := processRecord(record, uc); err != nil {
			uc.Logger.Err(err).Msgf("erro process record: %v", err)
			return err
		}
	}

	return nil
}

func processRecord(record []string, uc *Usecase) error {
	countryName := record[2]
	iso3 := record[1]

	dateStr := strings.TrimSpace(record[4])
	uc.Logger.Info().Msgf("Date: %s", dateStr)
	date, err := util.ParseDate(dateStr)
	if err != nil {
		uc.Logger.Err(err).Msgf("erro parse date record %v", err)

		return err
	}

	totalCases, err := util.ParseInt(record[5])
	if err != nil {
		uc.Logger.Err(err).Msgf("erro parse int record %s: %v", record[5], err)

		return err
	}

	totalDeaths, err := util.ParseInt(record[6])
	if err != nil {
		uc.Logger.Err(err).Msgf("erro parse int record %s: %v", record[6], err)

		return err
	}

	totalVaccinated, err := util.ParseInt(record[9])
	if err != nil {
		uc.Logger.Err(err).Msgf("erro parse int record %s: %v", record[9], err)

		return err
	}

	vaccineName := record[11]

	uc.Usecase.CreateCountry(countryName, iso3)
	uc.Usecase.CreateVaccine(vaccineName, iso3)
	uc.Usecase.CreateCovidCase(iso3, date, totalCases, totalDeaths)
	uc.Usecase.CreateVaccinationStats(iso3, date, totalVaccinated)
	uc.Usecase.CreateVaccineApproval(iso3, countryName, date, vaccineName)

	return nil
}

func renameCSV(reader *csv.Reader, usecase *Usecase) {
	usecase.Logger.Info().Str("hanlder", usecase.Path).Msg("Rename CSV")
}

func processCSVIfNotDone(reader *csv.Reader, usecase *Usecase) {
	usecase.Logger.Info().Str("hanlder", usecase.Path).Msg("Process CSV If Not Done")
}
