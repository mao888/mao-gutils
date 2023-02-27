package main

import "reflect"

// JudgeType 判断接口具体类型
func JudgeType(v interface{}) string {
	switch i := v.(type) {
	case int:
		return "int"
	case int64:
		return "int64"
	case int32:
		return "int32"
	case string:
		return "string"
	case float64:
		return "float64"
	case nil:
		return ""
	case bool:
		return "bool"
	default:
		_ = i
		return "unknown"
	}
}

// JudgeTypeByReflect 使用反射判断接口具体类型
func JudgeTypeByReflect(v interface{}) string {
	return reflect.TypeOf(v).String()
}
