package gutil

import (
	"fmt"
	"testing"
)

func TestValid(t *testing.T) {
	type args struct {
		json []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Normal condition",
			args: args{json: []byte(`{"name":"1","age":10}`)},
			want: true,
		},
		{
			name: "Empty JSON",
			args: args{json: []byte(`{}`)},
			want: true,
		},
		{
			name: "Error condition, only string",
			args: args{json: []byte(`aaaa`)},
			want: false,
		},
		{
			name: "Error condition, JSON error",
			args: args{json: []byte(`{"name""1","age":10}`)},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Valid(tt.args.json); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHuffmanEncode(t *testing.T) {
	body := []byte(`{"name":"zhangsan","info":[{"name":"class-01",id:"01"},{"name":"class-02",id:"02"},{"name":"class-03",id:"03"},{"name":"class-04",id:"04"},{"name":"class-05",id:"05"},{"name":"class-06",id:"06"},{"name":"class-07",id:"07"},{"name":"class-08",id:"08"},{"name":"class-09",id:"09"},{"name":"class-10",id:"10"}]}`)
	fmt.Println("压缩前大小：", len(body))
	result := HuffmanEncode(body)
	fmt.Println("压缩后大小：", len(result))
	result = HuffmanDecode(result)
	fmt.Println("解压后数据：", string(result))
	fmt.Println("比较结果：", string(result) == string(body))
}

func TestGzipEncode(t *testing.T) {
	body := []byte(`{"name":"zhangsan","info":[{"name":"class-01",id:"01"},{"name":"class-02",id:"02"},{"name":"class-03",id:"03"},{"name":"class-04",id:"04"},{"name":"class-05",id:"05"},{"name":"class-06",id:"06"},{"name":"class-07",id:"07"},{"name":"class-08",id:"08"},{"name":"class-09",id:"09"},{"name":"class-10",id:"10"}]}`)
	fmt.Println("压缩前大小：", len(body))
	result := GzipEncode(body)
	fmt.Println("压缩后大小：", len(result))
	result = GzipDecode(result)
	fmt.Println("解压后数据：", string(result))
	fmt.Println("比较结果：", string(result) == string(body))
}
