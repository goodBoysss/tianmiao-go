// 结构体辅助方法
package helpers

import "encoding/json"

// JsonEncode json加密
func JsonEncode(data interface{}) string {
	jsons, err := json.Marshal(data)
	if err == nil {

	}
	return string(jsons)
}

// JsonDecode json加密
func JsonDecode(data string) map[string]interface{} {
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(data), &dat)
	if err == nil {

	}
	return dat
}
