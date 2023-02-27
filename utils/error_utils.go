package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CheckSimpleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CheckError(c *gin.Context, err error, status int) {
  if err != nil{
	c.JSON(status, gin.H{
		 "message" : err.Error(),
    })
	return 
  } 
}

func CheckValidCountryName(c *gin.Context, status int, message string) {
	// if  countryNameParam == " " {
    //    c.JSON(status, gin.H{
	// 	 "message" : "Param Field Is Empty",
    // })
	
	// } else if countryNameParam != data.get(countryNameParam){
	// 	 c.JSON(status, gin.H{
	// 	 "message" : "Country Name Not Found",
	// })

	c.JSON(status, gin.H{"message" : message})
}