package gutil

import "testing"

func TestDingTalkAlarm(t *testing.T) {
	type args struct {
		serverName string
		message    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{serverName: "实验平台", message: "清楚缓存失败，实验ID：123003"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DingTalkAlarm(tt.args.serverName, tt.args.message); got != tt.want {
				t.Errorf("DingTalkAlarm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDingTalkAlarmUrl(t *testing.T) {
	type args struct {
		url        string
		serverName string
		message    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DingTalkAlarmUrl(tt.args.url, tt.args.serverName, tt.args.message); got != tt.want {
				t.Errorf("DingTalkAlarmUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
