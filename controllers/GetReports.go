package controllers

import (
	"net/http"

	"github.com/HasnathJami/Covid-Data-Analyzer/models"
	"github.com/HasnathJami/Covid-Data-Analyzer/utils"
	"github.com/gin-gonic/gin"
)


func GetReports(c *gin.Context) {
   countries := models.ReadCountries()
   var report models.Report

   for _, value := range countries{
	  report.TotalCases += value.Cases
	  report.TotalDeaths += value.Deaths
	  report.TotalRecovered += value.Recovered
   }

   utils.ProcessSuccess(c, "Reports Have Been Fetched Successfully", http.StatusOK, report)

}