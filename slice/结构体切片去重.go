package slice

type SeriesRes struct {
	SeriesName        string  `json:"seriesName"`        // 集合名称
	SeriesId          string  `json:"seriesId"`          // 集合id
	ExternalUrl       string  `json:"externalUrl"`       // 集合封面图
	SeriesDescription string  `json:"seriesDescription"` // 描述
	Issuer            string  `json:"issuer"`            // 发行方
	DnaCounts         int64   `json:"dnaCounts"`         // 累计发行数量/dna数字资产数
	DnaPrice          float64 `json:"dnaPrice"`          // 发行金额
}

// RemoveRepByMap 通过map主键唯一的特性过滤重复元素
// 结构体切片去重
func removeStructRepByMap(slc []*SeriesRes) []*SeriesRes {
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
