package routes

import (
	"github.com/gin-gonic/gin"
	"tianmiao-go/routes/admin"
	"tianmiao-go/routes/api"
	"tianmiao-go/routes/redirect"
	"tianmiao-go/routes/web"
)

func RouteRegister(router *gin.Engine) {
	//redirect
	redirect.RedirectRoutes(router)
	//api
	api.ApiRoutes(router)
	//web
	web.WebRoutes(router)
	//admin
	admin.AdminRoutes(router)
}
