package gutil

import (
	"github.com/mao888/go-utils/constants"
	"strings"
)

// MapSplitByComma 根据 "," 分割map为map数组
func MapSplitByComma(fields map[string]string) []map[string]string {
	mapFields := make([]map[string]string, 0)
	// range map
	for k, m := range fields {
		// 根据 "," 分割map中value
		multipleValues := strings.Split(m, constants.Comma) // [1 2 3]

		if len(multipleValues) > 1 { // 如果分割出来的数组长度 > 1 ,说明value中有 ","
			// range 分割出来的数组
			for _, v := range multipleValues {
				fields[k] = v                        // 将数组中遍历出来的值 替换掉原key 的value
				mapResult := make(map[string]string) // 临时map
				// range fields 将原map的值装入临时map
				for kf, f := range fields {
					mapResult[kf] = f
				}
				mapFields = append(mapFields, mapResult) // 临时map装入最终map数组
			}
			break
		}
		// 没有",",直接装入终map数组
		mapFields = append(mapFields, fields)
		break
	}
	return mapFields
}
