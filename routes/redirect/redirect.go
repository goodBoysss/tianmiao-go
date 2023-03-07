package redirect

import (
	"github.com/gin-gonic/gin"
	"tianmiao-go/app/Http/Controllers/Redirect"
)

func RedirectRoutes(router *gin.Engine) {
	RedirectController := new(Redirect.RedirectCroller)
	router.GET("/:short_key", RedirectController.RedirectUrl)
}
