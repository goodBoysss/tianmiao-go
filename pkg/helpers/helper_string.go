// 字符串辅助方法
package helpers

import (
	"strings"
)

// 字符串下划线转驼峰
func UnderlineToCamel(str string) string {
	output := ""

	if str == "" || str == "_" {
		return output
	}

	arr := strings.Split(str, "_")
	for _, value := range arr {
		if value != "" {
			//首个字符串转大写
			output += strings.ToUpper(value[:1]) + value[1:]
		}
	}
	return output
}
