package Response

import (
	"github.com/gin-gonic/gin"
	"tianmiao-go/pkg/response"
)

func SendData(c *gin.Context, data interface{}, msg string) {
	var res = make(map[string]interface{}, 3)
	res["code"] = 200

	//默认空对象
	if data == nil {
		data = map[string]string{}
	}
	res["data"] = data

	//文字描述
	if msg == "" {
		msg = "SUCCESS"
	}
	res["msg"] = msg
	response.JSON(c, res)
}
