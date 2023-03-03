// 结构体辅助方法
package helpers

import (
	"fmt"
	"github.com/spf13/cast"
	"reflect"
)

// 切片数组提取int字段
func SliceArrayColumnInt(input []map[string]interface{}, column string) []int {
	output := make([]int, len(input))
	for i, info := range input {
		output[i] = cast.ToInt(info[column])
		fmt.Println(info[column])
		fmt.Println(reflect.TypeOf(info[column]))
	}
	return output
}
