package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	router.GET("/player-data", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"player": "demo-player",
			"level":  7,
			"status": "active",
		})
	})

	// Gin listens on 0.0.0.0:8080 by default.
	_ = router.Run()
}
