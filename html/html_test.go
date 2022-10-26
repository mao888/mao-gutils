package gutil

import "testing"

func TestTrimHtml(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "style",
			args: args{"<style src=''></style>style测试"},
			want: "style测试",
		},
		{
			name: "script",
			args: args{"<script src=''/>script测试"},
			want: "script测试",
		},
		{
			name: "div",
			args: args{"<div>测</div>试测试"},
			want: "测试测试",
		},
		{
			name: "div_attr",
			args: args{"<div id='id'>测</div>试测试"},
			want: "测试测试",
		},
		{
			name: "p",
			args: args{"<p><img src='https://gia-art-needs.s3.ap-northeast-1.amazonaws.com/art-needs/debug/2021-10-21/gia-art-need1634801528303.png' alt=''/>在如</p>今快节奏的生活环境下"},
			want: "在如今快节奏的生活环境下",
		},
		{
			name: "span",
			args: args{"<span style='font-size: 1em;'>健康</span>快乐"},
			want: "健康快乐",
		},
		{
			name: "br",
			args: args{"<br/>健康<br/>快乐"},
			want: "健康快乐",
		},
		{
			name: "none",
			args: args{"健康快乐[div,css,span,img,br]"},
			want: "健康快乐[div,css,span,img,br]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimHtml(tt.args.src); got != tt.want {
				t.Errorf("TrimHtml() = %v, want %v", got, tt.want)
			}
		})
	}
}
