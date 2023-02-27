package web

import (
	"github.com/gin-gonic/gin"
)

func WebRoutes(router *gin.Engine) {
	router.GET("/web/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
