package time

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFormatTimeIfNotZero(t *testing.T) {
	type args struct {
		time   time.Time
		layout string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, FormatTimeIfNotZero(tt.args.time, tt.args.layout), "FormatTimeIfNotZero(%v, %v)", tt.args.time, tt.args.layout)
		})
	}
}

func TestGetBetweenDates(t *testing.T) {
	type args struct {
		startTime time.Time
		endTime   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBetweenDates(tt.args.startTime, tt.args.endTime)
			if !tt.wantErr(t, err, fmt.Sprintf("GetBetweenDates(%v, %v)", tt.args.startTime, tt.args.endTime)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetBetweenDates(%v, %v)", tt.args.startTime, tt.args.endTime)
		})
	}
}

func TestGetBetweenMonths(t *testing.T) {
	type args struct {
		startTime time.Time
		endTime   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBetweenMonths(tt.args.startTime, tt.args.endTime)
			if !tt.wantErr(t, err, fmt.Sprintf("GetBetweenMonths(%v, %v)", tt.args.startTime, tt.args.endTime)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetBetweenMonths(%v, %v)", tt.args.startTime, tt.args.endTime)
		})
	}
}

func TestNowTime(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NowTime(), "NowTime()")
		})
	}
}

func TestNowUnix(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NowUnix(), "NowUnix()")
		})
	}
}

func TestUnixToFormatTime(t *testing.T) {
	type args struct {
		timeStamp int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, UnixToFormatTime(tt.args.timeStamp), "UnixToFormatTime(%v)", tt.args.timeStamp)
		})
	}
}
