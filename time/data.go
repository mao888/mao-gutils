package time

import "time"

//DateCCT 北京时区
//注意： 不能使用带有时区的格式化
func DateCST(layout string) string {
	now := time.Now().UTC().Add(time.Hour * 8)
	return now.Format(layout)
}

//DatePST 美国标准时区
//注意： 不能使用带有时区的格式化
func DatePST(layout string) string {
	now := time.Now().UTC().Add(time.Hour * -8)
	return now.Format(layout)
}

//DateJST 日本时区
//注意： 不能使用带有时区的格式化
func DateJST(layout string) string {
	now := time.Now().UTC().Add(time.Hour * 9)
	return now.Format(layout)
}

//DateKST 韩国时区
//注意： 不能使用带有时区的格式化
func DateKST(layout string) string {
	now := time.Now().UTC().Add(time.Hour * 9)
	return now.Format(layout)
}

//DateUTC UTC时间
//注意： 不能使用带有时区的格式化
func DateUTC(layout string) string {
	now := time.Now().UTC()
	return now.Format(layout)
}
