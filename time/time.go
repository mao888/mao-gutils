package time

import (
	"errors"
	"time"
)

// NowTime 获取当前时间
func NowTime() string {
	return time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
}

// NowUnix 获取当前时间戳
func NowUnix() int64 {
	return time.Now().Unix()
}

// UnixToFormatTime 获取当前时间
func UnixToFormatTime(timeStamp int64) string {
	return time.Unix(timeStamp, 0).Format("2006-01-02 15:04:05")
}

func FormatTimeIfNotZero(time time.Time, layout string) string {
	if time.IsZero() {
		return ""
	}
	return time.Format(layout)
}

// GetBetweenMonths 根据开始日期和结束日期计算出时间段内所有日期
func GetBetweenMonths(startTime, endTime time.Time) ([]string, error) {
	d := []string{}
	timeFormatTpl := "2006-01"
	start := startTime.Format(timeFormatTpl)
	end := endTime.Format(timeFormatTpl)

	if len(timeFormatTpl) != len(start) {
		timeFormatTpl = timeFormatTpl[0:len(start)]
	}
	date, err := time.Parse(timeFormatTpl, start)
	if err != nil {
		return nil, err
	}
	date2, err := time.Parse(timeFormatTpl, end)
	if err != nil {
		return nil, err
	}
	if date2.Before(date) {
		return nil, errors.New("end before start")
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 1, 0)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d, nil
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
func GetBetweenDates(startTime, endTime time.Time) ([]string, error) {
	d := []string{}
	timeFormatTpl := "2006-01-02"
	start := startTime.Format(timeFormatTpl)
	end := endTime.Format(timeFormatTpl)

	if len(timeFormatTpl) != len(start) {
		timeFormatTpl = timeFormatTpl[0:len(start)]
	}
	date, err := time.Parse(timeFormatTpl, start)
	if err != nil {
		return nil, err
	}
	date2, err := time.Parse(timeFormatTpl, end)
	if err != nil {
		return nil, err
	}
	if date2.Before(date) {
		return nil, errors.New("end before start")
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d, nil
}
