package Response

import (
	"github.com/gin-gonic/gin"
	"tianmiao-go/pkg/response"
)

func SendData(c *gin.Context, data interface{}, msg string) {
	var res = make(map[string]interface{}, 3)
	res["code"] = 200
	res["data"] = data
	if msg == "" {
		msg = "请求成功"
	}
	res["msg"] = msg
	response.JSON(c, res)
}
