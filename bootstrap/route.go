package bootstrap

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"tianmiao-go/pkg/config"
	"tianmiao-go/routes"
)

func InitRoute() {
	// gin 实例
	router := gin.Default()

	// 注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册路由
	routes.RouteRegister(router)

	// 配置 404 路由

	setup404Handler(router)

	//加载html模板
	loadTemplates(router)

	//运行
	runRouter(router)
}

// 全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
	//middlewares.Logger(),
	//middlewares.Recovery(),
	//middlewares.ForceUA(),
	)
}

// 404返回
func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}

// 加载html模板
func loadTemplates(router *gin.Engine) {
	router.LoadHTMLGlob("./app/Templates/Redirect/Cover.html")
}

func runRouter(router *gin.Engine) {
	// 运行服务器
	err := router.Run("0.0.0.0:" + config.GetString("APP_PORT", "80"))
	if err != nil {
		log.Fatal(err)
	}
}
