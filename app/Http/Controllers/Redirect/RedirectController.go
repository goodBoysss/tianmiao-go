package Redirect

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tianmiao-go/app/Logics"
	"tianmiao-go/app/Response"
	"tianmiao-go/pkg/helpers"
	"time"
)

type RedirectCroller struct{}

// RedirectUrl 重定向跳转url
/*
 * User: zhanglinxiao<zhanglinxiao@tianmtech.cn>
 * DateTime: 2023/03/03 15:19
 * ApiLink:get /{$shortKey}
 * @return string
 */
func (r *RedirectCroller) RedirectUrl(c *gin.Context) {
	shortKey := c.Param("short_key")
	domain := c.Request.Host
	userAgent := c.Request.UserAgent()
	ip := c.RemoteIP()

	//初始化RedirectLogic
	logic := Logics.RedirectLogic{}
	//获取跳转信息
	redirectInfo := logic.GetRedirectInfo(domain, shortKey)
	if redirectInfo["origin_url"] != nil {
		//goroute 协程处理添加日志信息，由于http常驻进程，单次请求无需等待处理完成
		go func(redirectInfo map[string]interface{}, domain string, shortKey string, userAgent string, ip string) {
			//添加访问记录到redis缓存
			visitRecord := make(map[string]interface{})
			visitRecord["app_id"] = redirectInfo["app_id"]
			visitRecord["domain"] = domain
			visitRecord["short_key"] = shortKey
			visitRecord["user_agent"] = userAgent
			visitRecord["ip"] = ip
			visitRecord["visit_time"] = helpers.DateTime("Y-m-d H:i:s", time.Now().Unix())
			//将访问记录添加到redis缓存
			logic.AddRedirectVisitRecordToCache(visitRecord)
		}(redirectInfo, domain, shortKey, userAgent, ip)

		//获取浏览器类型：0-其他；1-微信应用内置；2-QQ应用内置
		browserType := logic.GetBrowserType(userAgent)
		if browserType == 0 { //普通浏览器直接跳转
			//302跳转
			c.Header("Cache-Control", "no-cache")
			c.Redirect(http.StatusFound, fmt.Sprintf("%s", redirectInfo["origin_url"]))

		} else { //特殊浏览器：例如1-微信应用内置；2-QQ应用内置，展示过度页面
			c.HTML(http.StatusOK, "cover.html", gin.H{
				"info": redirectInfo,
			})
		}
		return
	} else {
		Response.SendData(c, nil, "链接已失效！")
	}
}
