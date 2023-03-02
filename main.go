package main

import (
	"fmt"
	"tianmiao-go/app/Models"
	"tianmiao-go/bootstrap"
	"tianmiao-go/pkg/database"
)

func main() {
	//应用初始化
	bootstrap.InitApp()
	// 初始化日志
	bootstrap.SetupLogger()
	// 初始化数据库
	bootstrap.SetupDB()

	repo := database.Repo{}

	where := [][]string{
		[]string{"app_id", "=", "1"},
		[]string{"id", "=", "92"},
	}

	columns := []string{
		"id",
	}
	orderBy := map[string]string{
		"id": "asc",
	}
	o := repo.Source(&Models.RedirectUrl{}).GetOne(where, columns, orderBy)
	fmt.Println(o)
	//路由初始化（监听端口）
	//bootstrap.InitRoute()
}
