// 结构体辅助方法
package helpers

import (
	"github.com/mitchellh/mapstructure"
)

// 结构体转Map集合
func StructToMap(obj interface{}) map[string]interface{} {
	// 将对象转化为 map 类型
	output := make(map[string]interface{})
	err := mapstructure.Decode(obj, &output)
	if err != nil {
		// 处理错误
		err.Error()
	}
	return output
}
