package Logics

import (
	"fmt"
	"regexp"
	"tianmiao-go/app/Enums"
	"tianmiao-go/app/Models"
	"tianmiao-go/pkg/database/repo"
	"tianmiao-go/pkg/helpers"
	"tianmiao-go/pkg/redis"
)

type RedirectLogic struct {
}

// GetRedirectInfo 获取重定向url信息
/*
 * User: zhanglinxiao<zhanglinxiao@tianmtech.cn>
 * DateTime: 2023/03/03 15:19
 */
func (l *RedirectLogic) GetRedirectInfo(domain string, shortKey string) map[string]interface{} {
	redirectInfo := make(map[string]interface{})
	//md5加密
	domainMd5 := helpers.Md5(domain)

	//先从缓存获取
	redisKey := fmt.Sprintf(Enums.REDIS_KEY_REDIRECT_URLS, domainMd5, shortKey)
	redisValue := redis.Redis.Get(redisKey)

	redirectInfo = helpers.JsonDecode(redisValue)
	if len(redirectInfo) == 0 {
		//重新赋值
		redirectInfo = make(map[string]interface{})
		//获取跳转信息
		redirectModel := Models.RedirectUrl{}
		where := [][]string{
			[]string{"domain_md5", "=", domainMd5},
			[]string{"BINARY short_key", "=", shortKey},
		}
		repo.GetOne(&redirectModel, where, []string{"app_id", "origin_url", "is_show_cover"}, nil)

		if redirectModel.AppId > 0 {
			//是否展示跳转封面
			isShowCover := redirectModel.IsShowCover
			//封面图片地址
			coverImageUrl := ""

			redirectInfo["app_id"] = redirectModel.AppId
			redirectInfo["origin_url"] = redirectModel.OriginUrl
			redirectInfo["is_show_cover"] = redirectModel.IsShowCover
			if isShowCover == 1 {
				redirectCoverModel := Models.RedirectCover{}
				repo.GetOne(&redirectCoverModel, [][]string{
					[]string{"domain_md5", "=", domainMd5},
					[]string{"BINARY short_key", "=", shortKey},
				}, []string{"cover_image_url"}, nil)
				if redirectCoverModel.CoverImageUrl != "" {
					coverImageUrl = redirectCoverModel.CoverImageUrl
				}
			}

			redirectInfo["cover_image_url"] = coverImageUrl
		}
	}

	return redirectInfo
}

// AddRedirectVisitRecordToCache 添加访问记录到redis缓存
/*
 * User: zhanglinxiao<zhanglinxiao@tianmtech.cn>
 * DateTime: 2023/03/03 15:19
 */
func (l *RedirectLogic) AddRedirectVisitRecordToCache(visitRecord map[string]interface{}) {
	key := Enums.REDIS_KEY_REDIRECT_VISIT_RECORD
	redis.Redis.Client.LPush(redis.Redis.Context, key, helpers.JsonEncode(visitRecord))
}

// GetBrowserType 获取浏览器类型：0-其他；1-微信应用内置；2-QQ应用内置
/*
 * User: zhanglinxiao<zhanglinxiao@tianmtech.cn>
 * DateTime: 2023/03/03 15:19
 */
func (l *RedirectLogic) GetBrowserType(userAgent string) int {
	//0-其他；1-微信应用内置；2-QQ应用内置
	browserType := 0

	if browserType == 0 {
		reWx := regexp.MustCompile("MicroMessenger")
		if reWx.FindString(userAgent) != "" {
			browserType = 1
		}
	}

	if browserType == 0 {
		reQQ := regexp.MustCompile("QQ/[0-9]*")
		if reQQ.FindString(userAgent) != "" {
			browserType = 2
		}
	}

	return browserType
}
