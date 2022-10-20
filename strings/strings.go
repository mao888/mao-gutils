package gutil

import (
	"strconv"
	"strings"
)

// TrimSpace 去除空格
func TrimSpace(str string) string {
	return strings.Replace(str, " ", "", -1)
}

//VersionOrdinal 返回可比较的字符串，当传入非法ASCII码时返回空字符串
//用于版本比较
func VersionOrdinal(version string) string {
	const maxByte = 1<<8 - 1
	vo := make([]byte, 0, len(version)+8)
	j := -1
	for i := 0; i < len(version); i++ {
		b := version[i]
		if '0' > b || b > '9' {
			vo = append(vo, b)
			j = -1
			continue
		}
		if j == -1 {
			vo = append(vo, 0)
			j = len(vo) - 1
		}
		if vo[j] == 1 && vo[j+1] == '0' {
			vo[j+1] = b
			continue
		}
		if vo[j]+1 > maxByte {
			return ""
		}
		vo = append(vo, b)
		vo[j]++
	}
	return string(vo)
}

//VersionGreater 比较两个版本号。版本号只有数字和点组成
// 如：versionA == versionB  返回 0
// 如：versionA > versionB  返回 1
// 如：versionA < versionB  返回 -1
func VersionGreater(versionA, versionB string) int {
	tempA := VersionOrdinal(versionA)
	tempB := VersionOrdinal(versionB)
	switch {
	case tempA == tempB:
		return 0
	case tempA > tempB:
		return 1
	default:
		return -1
	}
}

//StringJoin 通过Builder拼接字符串
func StringJoin(strs ...string) string {
	var baseKey strings.Builder
	for _, value := range strs {
		baseKey.WriteString(value)
	}
	return baseKey.String()
}

func InSplitString(s string, substr string) bool {
	if !strings.HasSuffix(s, ",") {
		s += ","
	}
	if !strings.HasSuffix(substr, ",") {
		substr += ","
	}
	return strings.Contains(s, substr)
}

//IsExactExist 精准的字符串匹配，区分大小写
func IsExactExist(array []string, row string) bool {
	for i := range array {
		if array[i] == row {
			return true
		}
	}
	return false
}

//Str2Int32Array 讲字符串切片转化成int32切片
func Str2Int32Array(strArray []string) ([]int32, error) {
	result := make([]int32, 0, len(strArray))
	for _, str := range strArray {
		intRow, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		result = append(result, int32(intRow))
	}
	return result, nil
}
