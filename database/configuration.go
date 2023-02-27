package database

import (
	error_utils "github.com/HasnathJami/Covid-Data-Analyzer/utils"

	go_gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var (
	database *go_gorm.DB
)


func Connect(){
	db, err := go_gorm.Open("mysql", "root:123456@/covid_analyzer?charset=utf8&parseTime=True&loc=Local")
	error_utils.CheckSimpleError(err)
	database = db
}

func GetDB(s interface{}) *go_gorm.DB{
	Connect()
	database.AutoMigrate(s)
	return database
}