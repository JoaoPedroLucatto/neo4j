package entity

type VaccineApprovalOutPut struct {
	VaccineName  string `json:"vaccine_name"`
	ApprovalDate string `json:"approval_date"`
}

type CountriesByVaccineOutPut struct {
	CountryName string `json:"name"`
	Iso3        string `json:"iso3"`
}

type GetVaccinesUsedByCountryOutPut struct {
	VaccineName string
}
