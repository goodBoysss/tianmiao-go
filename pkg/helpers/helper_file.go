// 文件辅助方法
package helpers

import "os"

// FileExists 文件是否存在
func FileExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
