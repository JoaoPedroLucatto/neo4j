package usecase

import (
	"neo4j-etl/internal/entity"
	"time"
)

func (usecase *Usecase) CreateCountry(countryName, iso3 string) (*entity.Country, error) {
	covidData, err := entity.NewCountry(countryName, iso3)
	if err != nil {
		return nil, err
	}

	err = usecase.Repository.LoadCountries(covidData)
	if err != nil {
		return nil, err
	}

	return covidData, nil
}

func (usecase *Usecase) CreateVaccine(vaccineName, iso3 string) (*entity.Vaccine, error) {
	vaccine, err := entity.NewVaccine(vaccineName, iso3)
	if err != nil {
		return nil, err
	}

	err = usecase.Repository.LoadVaccines(vaccine)
	if err != nil {
		return nil, err
	}

	return vaccine, nil
}

func (usecase *Usecase) CreateCovidCase(countryIso string, date time.Time, totalcases int, totalDeaths int) (*entity.CovidCase, error) {
	covidData, err := entity.NewCovidCase(countryIso, date, totalcases, totalDeaths)
	if err != nil {
		return nil, err
	}

	err = usecase.Repository.LoadCovidCases(covidData)
	if err != nil {
		return nil, err
	}

	return covidData, nil
}

func (usecase *Usecase) CreateVaccinationStats(countryIso string, date time.Time, totalVaccinated int) (*entity.VaccinationStats, error) {
	covidData, err := entity.NewVaccinationStats(countryIso, date, totalVaccinated)
	if err != nil {
		return nil, err
	}

	err = usecase.Repository.LoadVaccinationStats(covidData)
	if err != nil {
		return nil, err
	}

	return covidData, nil
}

func (usecase *Usecase) CreateVaccineApproval(countryIso string, countryName string, date time.Time, vaccineName string) (*entity.VaccineApproval, error) {
	covidData, err := entity.NewVaccineApproval(countryIso, countryName, date, vaccineName)
	if err != nil {
		return nil, err
	}

	err = usecase.Repository.LoadVaccineApprovals(covidData)
	if err != nil {
		return nil, err
	}

	return covidData, nil
}
