package db

import "neo4j-etl/internal/entity"

type Repository interface {
	LoadCountries(*entity.Country) error
	LoadVaccines(*entity.Vaccine) error
	LoadCovidCases(*entity.CovidCase) error
	LoadVaccinationStats(*entity.VaccinationStats) error
	LoadVaccineApprovals(*entity.VaccineApproval) error
}
