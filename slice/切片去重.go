package slice

// RemoveRepByMap 通过map主键唯一的特性过滤重复元素
// 结构体切片去重
func removeRepByMap(slc []*SeriesRes) []*SeriesRes {
	resultMap := make(map[string]*SeriesRes, len(slc))
	for _, v := range slc {
		resultMap[v.SeriesId] = v
	}
	var result []*SeriesRes
	for _, v := range resultMap {
		result = append(result, v)
	}
	return result
}
