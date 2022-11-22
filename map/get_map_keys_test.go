/**
    @author:Hasee
    @data:2022/11/22
    @note:
**/
package gutil

import (
	"fmt"
	"testing"
)

func TestGetMapKeys(t *testing.T) {
	// 初始化map
	m := map[int]int{}
	for i := 0; i < 10000; i++ {
		m[i] = i
	}

	keys := getMapKeys1(m)
	fmt.Println(keys)
}
