package usecase

import (
	"neo4j-api/internal/entity"
	"time"
)

func (usecase *Usecase) GetCovidSummary(country string, date time.Time) (*map[string]interface{}, error) {
	covidSummaty, err := usecase.Repository.GetCovidSummary(country, date)
	if err != nil {
		return nil, err
	}

	return covidSummaty, nil
}

func (usecase *Usecase) GetFirstDoseCount(country string, date time.Time) (*map[string]interface{}, error) {
	dosgeCount, err := usecase.Repository.GetFirstDoseCount(country, date)
	if err != nil {
		return nil, err
	}

	return dosgeCount, nil
}

func (usecase *Usecase) GetVaccinesUsedByCountry(country string) (*map[string]interface{}, error) {
	usedbyCountry, err := usecase.Repository.GetVaccinesUsedByCountry(country)
	if err != nil {
		return nil, err
	}

	return usedbyCountry, nil
}

func (usecase *Usecase) GetVaccineApprovalDates(country string) (*[]entity.VaccineApprovalOutPut, error) {
	approvalDateserr, err := usecase.Repository.GetVaccineApprovalDates(country)
	if err != nil {
		return nil, err
	}

	return approvalDateserr, nil
}

func (usecase *Usecase) GetCountriesByVaccine(vaccineName string) (*[]entity.CountriesByVaccineOutPut, error) {
	countriesByVaccine, err := usecase.Repository.GetCountriesByVaccine(vaccineName)
	if err != nil {
		return nil, err
	}

	return countriesByVaccine, nil
}
