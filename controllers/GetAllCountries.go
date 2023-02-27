package controllers

import (
	"net/http"

	"github.com/HasnathJami/Covid-Data-Analyzer/models"
	"github.com/HasnathJami/Covid-Data-Analyzer/utils"
	"github.com/gin-gonic/gin"
)

func GetAllCountries(c *gin.Context) {
   //var  _ , cancel = context.WithTimeout(context.Background(), 100 * time.Second)
   countries := models.ReadCountries()
   utils.ProcessSuccess(c, "Countries Have Been Fetched Successfully", http.StatusOK, countries)
   //defer cancel()


}