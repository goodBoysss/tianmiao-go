package bootstrap

import (
	"github.com/gin-gonic/gin"
	btsConig "tianmiao-go/config"
	"tianmiao-go/pkg/config"
	"tianmiao-go/pkg/helpers"
)

func InitApp() {

	// 配置初始化，依赖命令行 --env 参数
	config.InitConfig(".env")
	appEnv := config.GetString("APP_ENV")
	if appEnv != "" && helpers.FileExists(".env."+appEnv) == true {
		config.InitConfig(appEnv)
	}

	//触发加载 config 包的所有 init 函数
	btsConig.Initialize()
	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	if config.GetBool("app.debug") == true {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

}
