package slice

import (
	"fmt"
	"testing"
)

// 输出两个slice切片的差集
func TestDifferenceSet(t *testing.T) {
	leyangjun1 := []int{1, 3, 5, 6}
	leyangjun2 := []int{1, 3, 5}

	retDiff := DifferenceSet(leyangjun1, leyangjun2)
	fmt.Println(retDiff)

}