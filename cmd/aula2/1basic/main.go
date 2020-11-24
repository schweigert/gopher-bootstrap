package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// curl localhost:8080 | jq .
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"olá": "mundo!"})
	})

	// curl -X POST localhost:8080 | jq .
	engine.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"olá": "mundo com POST!"})
	})

	engine.Run()
}
