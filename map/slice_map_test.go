package gutil

import (
	"reflect"
	"testing"
)

func TestFilterMapByMap(t *testing.T) {
	type args struct {
		filter  map[string]struct{}
		pram    map[string]string
		isExist bool
	}
	tests := []struct {
		name      string
		args      args
		wantArray map[string]string
	}{
		{
			name:      "1",
			args:      args{filter: map[string]struct{}{"1": {}, "2": {}}, pram: map[string]string{"1": "1", "2": "2", "3": "3"}, isExist: true},
			wantArray: map[string]string{"3": "3"},
		},
		{
			name:      "2",
			args:      args{filter: map[string]struct{}{"1": {}, "2": {}}, pram: map[string]string{"1": "1", "2": "2", "3": "3"}, isExist: false},
			wantArray: map[string]string{"1": "1", "2": "2"},
		},
		{
			name:      "3",
			args:      args{filter: map[string]struct{}{"1": {}, "2": {}}, pram: map[string]string{"2": "2", "3": "3"}, isExist: false},
			wantArray: map[string]string{"2": "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArray := FilterMapByMap(tt.args.filter, tt.args.pram, tt.args.isExist); !reflect.DeepEqual(gotArray, tt.wantArray) {
				t.Errorf("FilterMapByMap() = %v, want %v", gotArray, tt.wantArray)
			}
		})
	}
}

func TestFilterSliceByMap(t *testing.T) {
	type args struct {
		filter  map[string]struct{}
		pram    []string
		isExist bool
	}
	tests := []struct {
		name      string
		args      args
		wantArray []string
	}{
		{
			name:      "1",
			args:      args{filter: map[string]struct{}{"1": {}, "2": {}}, pram: []string{"1", "2", "3"}, isExist: true},
			wantArray: []string{"3"},
		},
		{
			name:      "2",
			args:      args{filter: map[string]struct{}{"1": {}, "2": {}}, pram: []string{"1", "2", "3"}, isExist: false},
			wantArray: []string{"1", "2"},
		},
		{
			name:      "3",
			args:      args{filter: map[string]struct{}{"1": {}, "2": {}}, pram: []string{"2", "3"}, isExist: false},
			wantArray: []string{"2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotArray := FilterSliceByMap(tt.args.filter, tt.args.pram, tt.args.isExist); !reflect.DeepEqual(gotArray, tt.wantArray) {
				t.Errorf("FilterSliceByMap() = %v, want %v", gotArray, tt.wantArray)
			}
		})
	}
}

func TestUniqueArray(t *testing.T) {
	type args struct {
		m []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "test1", args: args{m: []string{"1", "2", "3", "3", "2", "1"}}, want: []string{"1", "2", "3"}},
		{name: "test2", args: args{m: []string{"1", "2", "3"}}, want: []string{"1", "2", "3"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqueArray(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniqueArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
