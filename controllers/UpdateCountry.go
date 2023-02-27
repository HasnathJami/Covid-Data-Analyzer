package controllers

import (
	"net/http"
	"strings"

	"github.com/HasnathJami/Covid-Data-Analyzer/models"
	"github.com/HasnathJami/Covid-Data-Analyzer/utils"
	"github.com/gin-gonic/gin"
)

func UpdateCountry(c *gin.Context) {
  countryName := c.Params.ByName("country_name")

  var  updatedCountry models.Country
  err := c.BindJSON(&updatedCountry)
  utils.CheckError(c, err, http.StatusBadRequest)

  var countryDetails, _ = models.ReadCountryByCountryName(countryName)
  updatedCountry.ID = countryDetails.ID

  if strings.ToLower(countryName) != strings.ToLower(updatedCountry.CountryName) {
	  utils.CheckValidCountryName(c,http.StatusBadRequest, "Country Name Not Found")
  } else{
	 _ = models.ModifyCountry(&updatedCountry)
     utils.ProcessSuccess(c, "Country Has Been Updated Successfully", http.StatusOK, nil)
  }
}
