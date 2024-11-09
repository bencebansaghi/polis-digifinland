package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	parseSurveysFolder()
	router = gin.Default()
	log.Println("Starting server on port 8080")
	initializeApiRoutes()
	initializeRoutes()
	router.Run(":8080")
}
