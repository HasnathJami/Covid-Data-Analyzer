package controllers

import (
	"net/http"

	"github.com/HasnathJami/Covid-Data-Analyzer/models"
	"github.com/HasnathJami/Covid-Data-Analyzer/utils"
	"github.com/gin-gonic/gin"
)

func CreateCountry(c *gin.Context) {
  //var  _ , cancel = context.WithTimeout(context.Background(), 100 * time.Second)
  var  country models.Country
  err := c.BindJSON(&country)
  utils.CheckError(c, err, http.StatusBadRequest)
  value := country.InsertCountry()
  utils.ProcessSuccess(c, "Country Has Been Created Successfully",http.StatusOK, value)
  //defer cancel()
  
  }

  
