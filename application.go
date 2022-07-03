package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Working Fine env port:"+os.Getenv("PORT"))
	})
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Working Fine")
	})
	router.Run(":" + port)
}
