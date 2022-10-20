package gutil

// map、slice 过滤

func FilterSliceByMap(filter map[string]struct{}, pram []string, isExist bool) (array []string) {
	if len(filter) == 0 || len(pram) == 0 {
		return pram
	}
	for _, v := range pram {
		_, ok := filter[v]
		if ok && isExist || !ok && !isExist {
			continue
		}
		array = append(array, v)
	}
	return array
}

func FilterMapByMap(filter map[string]struct{}, pram map[string]string, isExist bool) (array map[string]string) {
	if len(filter) == 0 || len(pram) == 0 {
		return pram
	}
	array = make(map[string]string)
	for k, v := range pram {
		_, ok := filter[k]
		if ok && isExist || !ok && !isExist {
			continue
		}
		array[k] = v
	}
	return array
}

func MergeMap(m1, m2 map[string]interface{}) map[string]interface{} {
	if len(m1) == 0 {
		return m2
	}
	if len(m2) == 0 {
		return m1
	}
	for k, v := range m1 {
		m2[k] = v
	}
	return m2
}

// UniqueArray 数组去重
func UniqueArray(m []string) []string {
	reArray := make([]string, 0)
	tempMap := make(map[string]struct{})
	for _, v := range m {
		if _, ok := tempMap[v]; ok {
			continue
		}
		tempMap[v] = struct{}{}
		reArray = append(reArray, v)
	}
	return reArray
}
