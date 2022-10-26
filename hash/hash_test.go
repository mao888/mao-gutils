package gutil

import "testing"

func TestHash64(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "name", args: args{data: "1234563435433"}, want: Hash64("1234563435433")},
		{name: "name", args: args{data: "1234563435433"}, want: Hash64("1234563435433")},
		{name: "name", args: args{data: "1234563435433"}, want: 18338806214697512634},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash64(tt.args.data); got != tt.want {
				t.Errorf("Hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHash64Byte(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hash64Byte(tt.args.data); got != tt.want {
				t.Errorf("HashByte() = %v, want %v", got, tt.want)
			}
		})
	}
}
