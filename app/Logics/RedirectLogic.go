package Logics

import (
	"tianmiao-go/app/Models"
	"tianmiao-go/pkg/database/repo"
)

type RedirectLogic struct {
}

// 获取重定向url信息
func (l *RedirectLogic) GetRedirectInfo(domain string, shortKey string) interface{} {
	redirectModel := Models.RedirectUrl{}

	repo.GetOne(&redirectModel, nil, nil, nil)

	return redirectModel
}

// 添加访问记录到redis缓存
func (l *RedirectLogic) AddRedirectVisitRecordToCache(visitRecord map[string]string) interface{} {
	return visitRecord
}
