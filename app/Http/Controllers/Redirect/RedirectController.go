package Redirect

import (
	"github.com/gin-gonic/gin"
	"tianmiao-go/app/Logics"
	"tianmiao-go/app/Response"
)

type RedirectCroller struct{}

func (r *RedirectCroller) RedirectUrl(c *gin.Context) {
	domain := c.Request.Host
	shortKey := c.Param("short_key")
	//ip := c.RemoteIP()
	//逻辑层
	logic := new(Logics.RedirectLogic)
	redirectInfo := logic.GetRedirectInfo(domain, shortKey)
	if redirectInfo != nil {

	}

	//添加访问记录到redis缓存
	visitRecord := map[string]string{
		"app_id": "2",
	}
	logic.AddRedirectVisitRecordToCache(visitRecord)
	Response.SendData(c, redirectInfo, "")
}
