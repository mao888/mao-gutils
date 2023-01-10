package version

import (
	"regexp"
)

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

//VersionCheck 验证版本，该方法只支持app市场版本的格式：主版本.此版本.修订版本.热更版本。（且可以使用0开头）
//理论最大版本号：999.999.999.999
func VersionCheck(v string) bool {
	reg, _ := regexp.Compile("^\\d{1,3}(\\.(\\d{1,3})){1,3}$")
	return reg.MatchString(v)
}

//VersionApp 通过传入的版本号获取app的市场版本。
//app的市场版本格式为：x.x.x
func VersionApp(v string) string {
	return VersionAppByCount(v, 3)
}

//VersionAppByCount 通过传入的版本号和count确定返回几位的版本号
//v：版本号
//count：需要返回的版本号位数
func VersionAppByCount(v string, count int) string {
	for i, b := range v {
		if b != '.' {
			continue
		}
		if count--; count == 0 {
			return v[0:i]
		}
	}
	return v
}
