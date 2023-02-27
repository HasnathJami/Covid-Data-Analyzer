package models

import (
	"github.com/HasnathJami/Covid-Data-Analyzer/database"
	"github.com/jinzhu/gorm"
)

type Country struct{
	gorm.Model
	//CountryId        uint    `json:"country_id`
	CountryName      string  `json:"country_name"`
	Population       uint    `json:"population"`
	Cases            uint    `json:"cases"`
	Deaths           uint    `json:"deaths"`
	Recovered        uint    `json:"recovered"`	
}

var (
	db *gorm.DB
)

func init(){
	db = database.GetDB(&Country{})
}

func (country *Country)InsertCountry() *Country{
	//db.NewRecord(country)
    db.Create(country)
	return country
}

 func ReadCountries() []Country{
    var countries []Country
	db.Find(&countries)

	return countries
 }

func ReadCountryByCountryName(countryName string) (*Country,*gorm.DB){
   var country Country
   db := db.Where("country_name=?", countryName).Find(&country)

   return &country, db
}

 func ModifyCountry(country *Country) (*gorm.DB){
   db := db.Save(&country)

   return db
 }

 func RemoveCountry(countryName string) (*gorm.DB, *Country){
   var country Country
   db := db.Where("country_name=?", countryName).Delete(&country)

   return db, &country
 }