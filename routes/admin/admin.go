package admin

import (
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	router.GET("admin/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
