package entity

import (
	"time"
)

type Vaccine struct {
	Name       string
	CountryIso string
}

type Country struct {
	Name string `csv:"country"`
	Iso3 string `csv:"iso3"`
}

type CovidCase struct {
	CountryIso  string    `csv:"iso3"`
	Date        time.Time `csv:"date"`
	TotalCases  int       `csv:"total_cases"`
	TotalDeaths int       `csv:"total_deaths"`
}

type VaccinationStats struct {
	CountryIso      string    `csv:"iso3"`
	Date            time.Time `csv:"date"`
	TotalVaccinated int       `csv:"total_vaccinated"`
}

type VaccineApproval struct {
	CountryIso  string `csv:"iso3"`
	CountryName string
	Date        time.Time `csv:"date"`
	Vaccine     string    `csv:"vaccine"`
}

func NewCountry(countryName, iso3 string) (*Country, error) {
	return &Country{Name: countryName, Iso3: iso3}, nil
}

func NewVaccine(vaccineName, countryIso string) (*Vaccine, error) {
	return &Vaccine{Name: vaccineName, CountryIso: countryIso}, nil
}

func NewCovidCase(countryIso string, date time.Time, totalcases int, totalDeaths int) (*CovidCase, error) {
	return &CovidCase{CountryIso: countryIso, Date: date, TotalCases: totalcases, TotalDeaths: totalDeaths}, nil
}

func NewVaccinationStats(countryIso string, date time.Time, totalVaccinated int) (*VaccinationStats, error) {
	return &VaccinationStats{CountryIso: countryIso, Date: date, TotalVaccinated: totalVaccinated}, nil
}

func NewVaccineApproval(iso3 string, countryName string, date time.Time, vaccineName string) (*VaccineApproval, error) {
	return &VaccineApproval{CountryIso: iso3, CountryName: countryName, Date: date, Vaccine: vaccineName}, nil
}
