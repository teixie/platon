package helpers

import (
	"fmt"
	"strings"
)

// 从map取值，key支持a.b.c的形式，key两边的"."会自动踢除
func MapGet(m interface{}, k string) interface{} {
	k = strings.Trim(k, ".")
	if m == nil || k == "" {
		return nil
	}

	k1 := "" // 当前key
	k2 := "" // 子级key
	p := strings.Index(k, ".")
	if p < 0 { // 不存在"."符
		k1 = k
	} else {
		k1 = k[0:p]
		k2 = k[p+1:]
	}

	switch m.(type) {
	case map[string]interface{}:
		tmp := m.(map[string]interface{})
		if val, ok := tmp[k1]; ok && k2 == "" {
			return val
		} else if ok {
			return MapGet(val, k2)
		}
	case map[string]int:
		tmp := m.(map[string]int)
		if val, ok := tmp[k1]; ok && k2 == "" {
			return val
		}
	case map[string]int64:
		tmp := m.(map[string]int64)
		if val, ok := tmp[k1]; ok && k2 == "" {
			return val
		}
	case map[string]string:
		tmp := m.(map[string]string)
		if val, ok := tmp[k1]; ok && k2 == "" {
			return val
		}
	case map[string]bool:
		tmp := m.(map[string]bool)
		if val, ok := tmp[k1]; ok && k2 == "" {
			return val
		}
	}

	return nil
}

// 从map取值，key支持a.b.c的形式，key两边的"."会自动踢除
func MapGetInt64(m interface{}, k string) int64 {
	return ParseInt64(MapGet(m, k))
}

// 从map取值，key支持a.b.c的形式，key两边的"."会自动踢除
func MapGetString(m interface{}, k string) string {
	return fmt.Sprint(MapGet(m, k))
}
