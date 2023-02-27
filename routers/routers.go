package routers

import (
	"github.com/HasnathJami/Covid-Data-Analyzer/controllers"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine{
	router := gin.New()

	api := router.Group("/api")

	api.POST("/country", controllers.CreateCountry)
	api.GET("/countries", controllers.GetAllCountries)
	api.GET("/country/:country_name", controllers.GetCountryByCountryName)
	api.PUT("/country/:country_name", controllers.UpdateCountry)
	api.DELETE("/country/:country_name", controllers.DeleteCountry)

	api.GET("/reports", controllers.GetReports)

	return router
}