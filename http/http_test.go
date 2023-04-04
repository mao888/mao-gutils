package gutil

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestHttpGet(t *testing.T) {
	type args struct {
		url    string
		body   io.Reader
		header map[string][]string
	}
	tests := []struct {
		name           string
		args           args
		wantHttpStatus int
		wantResp       []byte
		wantErr        bool
	}{
		{
			name:           "HttpGet",
			args:           args{"https://internal-prophet.akgoo.net/dreamisland/heartbeat", nil, nil},
			wantHttpStatus: 200,
			wantResp:       make([]byte, 0),
			wantErr:        false,
		},
		{
			name:           "HttpGetHaveResp",
			args:           args{"https://internal-prophet.akgoo.net/dreamisland/version", nil, nil},
			wantHttpStatus: 200,
			wantResp:       []byte("{\"version\":\"prod-prophet-22-7c4f94669-gtxl9\"}"),
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHttpStatus, gotResp, err := HttpGet(tt.args.url, tt.args.body, tt.args.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("HttpGet() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("HttpGet() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestHttpPost(t *testing.T) {
	type args struct {
		url    string
		body   io.Reader
		header map[string][]string
	}
	tests := []struct {
		name           string
		args           args
		wantHttpStatus int
		wantResp       []byte
		wantErr        bool
	}{
		{
			name:           "HttpPost",
			args:           args{"https://internal-prophet.akgoo.net/dreamisland/heartbeat", nil, nil},
			wantHttpStatus: 404,
			wantResp:       nil,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHttpStatus, gotResp, err := HttpPost(tt.args.url, tt.args.body, tt.args.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("HttpPost() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("HttpPost() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestHttpPostForm(t *testing.T) {
	type args struct {
		postUrl string
		body    map[string][]string
		header  map[string][]string
	}
	tests := []struct {
		name           string
		args           args
		wantHttpStatus int
		wantResp       []byte
		wantErr        bool
	}{
		{
			name:           "HttpPost",
			args:           args{"https://internal-prophet.akgoo.net/dreamisland/heartbeat", nil, nil},
			wantHttpStatus: 404,
			wantResp:       nil,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHttpStatus, gotResp, err := HttpPostForm(tt.args.postUrl, tt.args.body, tt.args.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpPostForm() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("HttpPostForm() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("HttpPostForm() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestHttpPostJson(t *testing.T) {
	type args struct {
		url    string
		body   []byte
		header map[string][]string
	}
	tests := []struct {
		name           string
		args           args
		wantHttpStatus int
		wantResp       []byte
		wantErr        bool
	}{
		{
			name:           "PostJson",
			args:           args{"https://internal-prophet.akgoo.net/dreamisland/heartbeat", nil, nil},
			wantHttpStatus: 404,
			wantResp:       nil,
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHttpStatus, gotResp, err := HttpPostJson(tt.args.url, tt.args.body, tt.args.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpPostJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("HttpPostJson() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("HttpPostJson() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestHttpGetJson(t *testing.T) {
	type args struct {
		url    string
		body   []byte
		header map[string][]string
	}
	tests := []struct {
		name           string
		args           args
		wantHttpStatus int
		wantResp       []byte
		wantErr        bool
	}{
		{
			name:           "GetJson",
			args:           args{"https://internal-prophet.akgoo.net/dreamisland/heartbeat", nil, nil},
			wantHttpStatus: 200,
			wantResp:       make([]byte, 0),
			wantErr:        false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHttpStatus, gotResp, err := HttpGetJson(tt.args.url, tt.args.body, tt.args.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpGetJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHttpStatus != tt.wantHttpStatus {
				t.Errorf("HttpGetJson() gotHttpStatus = %v, want %v", gotHttpStatus, tt.wantHttpStatus)
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("HttpGetJson() gotResp = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}

func TestHttpPostMultipart(t *testing.T) {
	file := getFileBytes("https://ark-oss.bettagames.com/2023-03/9744ac8f667b20048590f0051b15e90d.mp4")
	fmt.Println("文件大小：", len(file))
	url := "https://ad.oceanengine.com/open_api/2/file/video/ad/"
	formData := map[string]string{
		"advertiser_id":   "1760312309087432",
		"upload_type":     "UPLOAD_BY_FILE",
		"video_signature": "9744ac8f667b20048590f0051b15e90d",
	}
	fileData := map[string]FileObject{"video_file": {
		Name:    "auto4_1111111111.11111111_游戏35和36__V_ZJR_ZJR_en_16X9_33s",
		Content: file,
	}}
	header := map[string]string{
		"Content-Type": "multipart/form-data",
		"Access-Token": "b6d470f1a2190665f6bb0d77e395911bb7384abf",
	}
	code, resp, err := HttpPostMultipart(url, formData, fileData, header)
	fmt.Println("响应状态码：", code)
	fmt.Println("响应体：", string(resp))
	fmt.Println(err)
}

func getFileBytes(url string) []byte {
	code, body, err := HttpGet(url, nil, nil)
	fmt.Println(code, err)
	return body
}
