package controllers

import (
	"net/http"
	"strings"

	"github.com/HasnathJami/Covid-Data-Analyzer/models"
	"github.com/HasnathJami/Covid-Data-Analyzer/utils"
	"github.com/gin-gonic/gin"
)

func GetCountryByCountryName(c *gin.Context) {
  countryName := c.Params.ByName("country_name")

  countryDetails, _ := models.ReadCountryByCountryName(countryName)
  if strings.ToLower(countryName) != strings.ToLower(countryDetails.CountryName) {
	  utils.CheckValidCountryName(c,http.StatusBadRequest, "Country Name Not Found")
      return
  }
  utils.ProcessSuccess(c, "Country Has Been Fetched By Country Name", http.StatusOK, countryDetails)

}