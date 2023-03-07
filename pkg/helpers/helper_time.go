// 结构体辅助方法
package helpers

import (
	"strings"
	"time"
)

// 结构体转Map集合
func DateTime(format string, t int64) string {
	format = strings.Replace(format, "Y", "2006", -1)
	format = strings.Replace(format, "y", "06", -1)
	format = strings.Replace(format, "m", "01", -1)
	format = strings.Replace(format, "d", "02", -1)
	format = strings.Replace(format, "H", "15", -1)
	format = strings.Replace(format, "h", "15", -1)
	format = strings.Replace(format, "i", "04", -1)
	format = strings.Replace(format, "s", "05", -1)
	return time.Unix(t, 0).Format(format)
}
