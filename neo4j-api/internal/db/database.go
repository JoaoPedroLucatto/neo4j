package db

import (
	"neo4j-api/internal/entity"
	"time"
)

type Repository interface {
	GetCovidSummary(string, time.Time) (*map[string]interface{}, error)
	GetFirstDoseCount(string, time.Time) (*map[string]interface{}, error)
	GetVaccinesUsedByCountry(string) (*map[string]interface{}, error)
	GetVaccineApprovalDates(string) (*[]entity.VaccineApprovalOutPut, error)
	GetCountriesByVaccine(string) (*[]entity.CountriesByVaccineOutPut, error)
}
