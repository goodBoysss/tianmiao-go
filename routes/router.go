package routes

import (
	"github.com/gin-gonic/gin"
	"tm/routes/admin"
	"tm/routes/api"
	"tm/routes/web"
)

func RouteRegister(router *gin.Engine) {
	//api
	api.ApiRoutes(router)
	//web
	web.WebRoutes(router)
	//admin
	admin.AdminRoutes(router)
}
