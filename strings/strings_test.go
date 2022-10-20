package gutil

import "testing"

func TestStringJoin(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "String Join",
			args: args{strs: []string{"h", "e", "l", "lo"}},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringJoin(tt.args.strs...); got != tt.want {
				t.Errorf("StringJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func TestIsExactExist(t *testing.T) {
	type args struct {
		array []string
		row   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				array: []string{"abc", "Abc", "ad"},
				row:   "ABC",
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				array: []string{"abc", "Abc", "ad"},
				row:   "abc",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				array: []string{"/abc", "Abc", "/ad"},
				row:   "/ad",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsExactExist(tt.args.array, tt.args.row); got != tt.want {
				t.Errorf("IsExactExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
