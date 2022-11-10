package os

import (
	gutil "github.com/mao888/go-utils/map"
	"os"
)

// ExpandByMap 替换template模版中 $ 符号后的字符串
func ExpandByMap(template string, fields map[string]string) string {
	t := os.Expand(template, func(k string) string { return fields[k] })
	return t
}

// ExpandByFun 自定义函数替换字符串模版中 $ 符号后的字符串
func ExpandByFun(str string, f func(string) string) string {
	return os.Expand(str, f)
}

/**
 *	输入示例：
 *	"command_fields": {
 *      "user_id": "1,2,3",
 *      "prop": "huChao"
 *   }
 *
 *  template: "chartid=${user_id}&prop=${prop}"
 *
 *  输出示例：
 *  [
 *       "chartid=1&prop=hudaoju",
 *       "chartid=2&prop=hudaoju",
 *       "chartid=3&prop=hudaoju"
 *   ]
 *
 */

// GetComposedTemplateListExpandByMap 使用替换template模版中 $ 符号后的字符串
// isMultiple: 是否根据 "," 分割
func GetComposedTemplateListExpandByMap(template string, isMultiple bool, fields map[string]string) []string {
	mapFields := make([]map[string]string, 0)

	if !isMultiple {
		return []string{ExpandByMap(template, fields)}
	}

	// 根据 "," 分割map为map数组
	mapFields = gutil.MapSplitByComma(fields)

	result := make([]string, 0, len(mapFields))
	for _, m := range mapFields {
		result = append(result, ExpandByMap(template, m))
	}
	return result
}
