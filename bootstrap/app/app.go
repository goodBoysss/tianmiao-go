package app

import (
	"github.com/gin-gonic/gin"
	"tm/pkg/config"
)

func Init() {

	// 配置初始化，依赖命令行 --env 参数
	config.InitConfig("../../.env")

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.DebugMode)

}
