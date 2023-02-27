package api

import (
	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
