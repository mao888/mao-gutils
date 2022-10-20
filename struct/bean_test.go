package _struct

import (
	"testing"
)

type Src1 struct {
	ID   int
	Name *string
	Addr string
}

type Src2 struct {
	Name *string
	Addr string
}

type Src3 struct {
	Name *string
	addr string
	Age  int
}

type Dst1 struct {
	ID   int
	Name *string
	Addr string
}

func TestBeanCopy(t *testing.T) {
	name := "zhangsan"
	type args struct {
		src interface{}
		dst interface{}
	}
	tests := []struct {
		name    string
		args    args
		re      interface{}
		wantErr bool
	}{
		{
			name: "BeanCopy",
			args: args{
				src: Src1{ID: 1, Name: &name, Addr: "北京"},
				dst: &Dst1{},
			},
			re:      Dst1{ID: 1, Name: &name, Addr: "北京"},
			wantErr: false,
		},
		{
			name: "NotInSrc",
			args: args{
				src: Src2{Name: &name, Addr: "北京"},
				dst: &Dst1{},
			},
			re:      Dst1{Name: &name, Addr: "北京"},
			wantErr: false,
		},
		{
			name: "NotInDst",
			args: args{
				src: Src3{Name: &name, addr: "北京", Age: 10},
				dst: &Dst1{},
			},
			re:      Dst1{Name: &name},
			wantErr: false,
		},
		{
			name: "SrcIsPtr",
			args: args{
				src: &Src3{Name: &name, addr: "北京", Age: 10},
				dst: &Dst1{},
			},
			re:      Dst1{Name: &name},
			wantErr: false,
		},
		{
			name: "DstNoPtr",
			args: args{
				src: Src3{Name: &name, addr: "北京", Age: 10},
				dst: Dst1{},
			},
			re:      Dst1{Name: &name},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := StructCopy(tt.args.src, tt.args.dst); (err != nil) != tt.wantErr {
				t.Errorf("BeanCopy() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				dst := tt.args.dst.(*Dst1)
				re := tt.re.(Dst1)
				if *dst != re {
					t.Errorf("BeanCopy() return = %v, want %v", dst, re)
				}
			}
		})
	}
}
