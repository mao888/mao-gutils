package os

import "os"

// ExpandByMap 替换template模版中 $ 符号后的字符串
func ExpandByMap(template string, fields map[string]string) string {
	t := os.Expand(template, func(k string) string { return fields[k] })
	return t
}

// ExpandByFun 自定义函数替换字符串模版中 $ 符号后的字符串
func ExpandByFun(str string, f func(string) string) string {
	return os.Expand(str, f)
}
