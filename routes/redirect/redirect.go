package redirect

import (
	"github.com/gin-gonic/gin"
	Redirect2 "tianmiao-go/app/Http/Controllers/Redirect"
)

func RedirectRoutes(router *gin.Engine) {
	RedirectController := new(Redirect2.RedirectCroller)
	router.GET("/:short_key", RedirectController.RedirectUrl)
}
