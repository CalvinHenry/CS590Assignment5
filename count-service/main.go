package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var count int = 0
func main() {
	router := gin.Default()
	router.GET("/requestcount", requestCount)

	router.Run(":8082")
}

func requestCount(c *gin.Context) {
	count = count + 1
	c.JSON(http.StatusOK, gin.H{
		"requestsMade": count,
	},
	)
}
