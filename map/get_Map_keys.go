package gutil

import "reflect"

// golang 获取map所有key的方式
//	方法1（效率很高）：
func getMapKeys1(m map[int]int) []int {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

//	方法2（效率很高）：
func getMapKeys2(m map[int]int) []int {
	// 数组默认长度为map长度,后面append时,不需要重新申请内存和拷贝,效率很高
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

//	方法3（效率较低）：
func getMapKeys3(m map[int]int) []int {
	// 注意：由于数组默认长度为0，后面append时，需要重新申请内存和拷贝，所以效率较低
	keys := []int{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

//	方法4（效率极低）：
func getMapKeys4(m map[int]int) int {
	// 注意:虽然此写法简洁,但MapKeys函数内部操作复杂,效率极低
	keys := reflect.ValueOf(m).MapKeys()
	return len(keys)
}
