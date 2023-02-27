package controllers

import (
	"net/http"

	"github.com/HasnathJami/Covid-Data-Analyzer/models"
	"github.com/HasnathJami/Covid-Data-Analyzer/utils"
	"github.com/gin-gonic/gin"
)

func DeleteCountry(c *gin.Context){
  countryName := c.Params.ByName("country_name")

  _, countryDetails := models.RemoveCountry(countryName)
  if countryDetails.CountryName != "" {
	  utils.CheckValidCountryName(c,http.StatusBadRequest, "Country Name Not Found")
    return
  }
  utils.ProcessSuccess(c, "Country Has Been Deleted Successfully", http.StatusOK, nil)
}