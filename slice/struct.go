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
