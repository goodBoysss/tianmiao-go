package Redirect

import (
	"github.com/gin-gonic/gin"
	"tianmiao-go/app/Logics"
	"tianmiao-go/pkg/response"
)

type RedirectCroller struct{}

func (r *RedirectCroller) RedirectUrl(c *gin.Context) {
	domain := c.Request.Host
	shortKey := c.Param("short_key")
	//ip := c.RemoteIP()
	//逻辑层
	logic := new(Logics.RedirectLogic)
	redirectInfo := logic.GetRedirectInfo(domain, shortKey)

	//添加访问记录到redis缓存
	visitRecord := map[string]string{
		"app_id": "2",
	}

	//'app_id' => $redirectInfo['app_id'] ?? 0,
	//	'domain' => $domain,
	//	'short_key' => $shortKey,
	//	'user_agent' => $userAgent,
	//	'ip' => $ip,
	//	'visit_time' => date("Y-m-d H:i:s"),

	logic.AddRedirectVisitRecordToCache(visitRecord)

	response.JSON(c, redirectInfo)
}
