package Logics

type RedirectLogic struct {
}

// 获取重定向url信息
func (l *RedirectLogic) GetRedirectInfo(domain string, shortKey string) map[string]interface{} {
	redirectInfo := map[string]interface{}{
		"app_id":          1,
		"origin_url":      "",
		"is_show_cover":   "",
		"cover_image_url": "",
	}

	//repo := database.Repo{}
	//fmt.Println(repo.Source(&Models.App{}).GetOne())

	return redirectInfo
}

// 添加访问记录到redis缓存
func (l *RedirectLogic) AddRedirectVisitRecordToCache(visitRecord map[string]string) interface{} {
	return visitRecord
}
