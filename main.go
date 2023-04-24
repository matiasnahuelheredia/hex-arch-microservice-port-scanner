package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.POST("/scanports", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "not Yet Implemented",
		})
	})

	// Iniciar el servidor HTTP
	router.Run(":8080")
}
