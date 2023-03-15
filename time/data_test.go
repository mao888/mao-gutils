package time

import "testing"

func TestDateCST(t *testing.T) {
	type args struct {
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
			if got := DateCST(tt.args.layout); got != tt.want {
				t.Errorf("DateCST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateJST(t *testing.T) {
	type args struct {
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
			if got := DateJST(tt.args.layout); got != tt.want {
				t.Errorf("DateJST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateKST(t *testing.T) {
	type args struct {
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
			if got := DateKST(tt.args.layout); got != tt.want {
				t.Errorf("DateKST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDatePST(t *testing.T) {
	type args struct {
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
			if got := DatePST(tt.args.layout); got != tt.want {
				t.Errorf("DatePST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDateUTC(t *testing.T) {
	type args struct {
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
			if got := DateUTC(tt.args.layout); got != tt.want {
				t.Errorf("DateUTC() = %v, want %v", got, tt.want)
			}
		})
	}
}
