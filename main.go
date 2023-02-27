package main

import (
	"tm/bootstrap/app"
	"tm/bootstrap/route"
)

func main() {
	//应用初始化
	app.Init()
	//路由初始化
	route.Init()
}
