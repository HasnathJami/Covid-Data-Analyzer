package main

import (
	"os"

	"github.com/HasnathJami/Covid-Data-Analyzer/routers"
	"github.com/HasnathJami/Covid-Data-Analyzer/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load(".env")
    utils.CheckSimpleError(err)

    var router *gin.Engine
    router = routers.Router()

    router.Use(gin.Logger())
    router.Use(cors.Default())
	
    router.Run(":" + os.Getenv("PORT"))
}
