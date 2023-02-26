package routers

import (
	"github.com/HasnathJami/Covid-Data-Analyzer/controllers"
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	router = gin.Default()

	router.Group("api")

	router.POST("/country", controllers.CreateCountry)
	router.GET("/countries", controllers.GetAllCountries)
	router.GET("/country/:country_name", controllers.GetCountryByCountryName)
	router.PUT("/country/:country_name", controllers.UpdateCountry)
	router.DELETE("/country/:country_name", controllers.DeleteCountry)

	router.GET("/reports", controllers.GetReports)
}