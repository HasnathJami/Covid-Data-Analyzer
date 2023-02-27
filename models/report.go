package models

import (
	"github.com/jinzhu/gorm"
)

type Report struct {
	gorm.Model
	ReportId   uint `json:"report_id"`
	TotalCases uint  `json:"total_cases"`
	TotalDeaths uint `json:"total_deaths"`
	TotalRecovered uint `json:"total_recovered"`
	//DeathRatePerPopulation string `json:"death_rate_per_population"` //// country-wise list
	//RecoveryRateWithTotalCases string `json:"recovery_rate_with_total_cases"` // country-wise list
	//SuccessfulCountry //countrywise list
}