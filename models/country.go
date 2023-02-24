pacakge models

import "gorm.io/gorm"

type Country struct{
	gorm.Model
	CountryId        uint    `json:"country_id`
	CountryName      string  `json:"country_name"`
	Population       uint    `json:"population"`
	Cases            uint    `json:"cases"`
	Deaths           uint    `json:"deaths"`
	Recovered        uint    `json:"recovered"`	
}