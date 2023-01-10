package version

import "testing"

func TestVersionOrdinal(t *testing.T) {
	type args struct {
		version string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "version1",
			args: args{version: "1.3.5"},
			want: "\u00011.\u00013.\u00015",
		},
		{
			name: "version2",
			args: args{version: "1.21.5"},
			want: "\u00011.\u000221.\u00015",
		},
		{
			name: "version3",
			args: args{version: "1.01.5"},
			want: "\u00011.\u00011.\u00015",
		},
		{
			name: "version3",
			args: args{version: "1.101.5"},
			want: "\u00011.\u0003101.\u00015",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VersionOrdinal(tt.args.version); got != tt.want {
				t.Errorf("VersionOrdinal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersionGreater(t *testing.T) {
	type args struct {
		versionA string
		versionB string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1.3.4 > 1.3.3",
			args: args{versionA: "1.3.4", versionB: "1.3.3"},
			want: 1,
		},
		{
			name: "1.10.3 > 1.9.4",
			args: args{versionA: "1.10.3", versionB: "1.9.4"},
			want: 1,
		},
		{
			name: "1.9.10 > 1.9.4",
			args: args{versionA: "1.9.10", versionB: "1.9.4"},
			want: 1,
		},
		{
			name: "1.9.4 == 1.9.4",
			args: args{versionA: "1.9.4", versionB: "1.9.4"},
			want: 0,
		},
		{
			name: "1.9.4 == 1.09.04",
			args: args{versionA: "1.9.4", versionB: "1.09.04"},
			want: 0,
		},
		{
			name: "1.3.3 < 1.3.4",
			args: args{versionA: "1.3.3", versionB: "1.3.4"},
			want: -1,
		},
		{
			name: "1.9.4 < 1.10.3",
			args: args{versionA: "1.9.4", versionB: "1.10.3"},
			want: -1,
		},
		{
			name: "1.9.4 > 1.09.10",
			args: args{versionA: "1.9.4", versionB: "1.09.10"},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VersionGreater(tt.args.versionA, tt.args.versionB); got != tt.want {
				t.Errorf("VersionGreater() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersionCheck(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "超出4位及超出999版本号:",
			args: args{v: "999.122.999.2.3"},
			want: false,
		},
		{
			name: "超出999版本号:",
			args: args{v: "999.122.1999"},
			want: false,
		},
		{
			name: "超出4位版本号:",
			args: args{v: "999.122.999.2.3"},
			want: false,
		},
		{
			name: "2位版本号:",
			args: args{v: "999.122"},
			want: true,
		},
		{
			name: "3位版本号:",
			args: args{v: "0.022.1"},
			want: true,
		},
		{
			name: "4位版本号:",
			args: args{v: "0.022.1.02"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VersionCheck(tt.args.v); got != tt.want {
				t.Errorf("VersionCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersionAppByCount1(t *testing.T) {
	type args struct {
		v     string
		count int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "四位版本号获取市场版本：",
			args: args{v: "10.0.3.4", count: 3},
			want: "10.0.3",
		},
		{
			name: "两位版本号获取市场版本：",
			args: args{v: "10.0", count: 3},
			want: "10.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VersionAppByCount(tt.args.v, tt.args.count); got != tt.want {
				t.Errorf("VersionAppByCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
