package helpers

import (
	"strconv"
)

// 转int64
func ParseInt64(val interface{}) int64 {
	if val == nil {
		return 0
	}
	switch val.(type) {
	case bool:
		if val.(bool) {
			return 1
		}
		return 0
	case int:
		return int64(val.(int))
	case int64:
		return val.(int64)
	case float64:
		return int64(val.(float64))
	case string:
		str := val.(string)
		num, _ := strconv.ParseInt(str, 10, 64)
		return num
	}
	return 0
}

// 转int
func ParseInt(val interface{}) int {
	return int(ParseInt64(val))
}
