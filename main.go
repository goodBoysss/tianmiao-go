package main

import (
	"tianmiao-go/bootstrap"
)

func main() {
	//应用初始化
	bootstrap.InitApp()
	// 初始化日志
	bootstrap.SetupLogger()
	// 初始化数据库
	bootstrap.SetupDB()
	//路由初始化（监听端口）
	bootstrap.InitRoute()
}
