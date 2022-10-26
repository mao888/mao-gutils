package gutil

import (
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
