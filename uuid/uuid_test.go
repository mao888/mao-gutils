package uuid

import (
	"fmt"
	"testing"
)

func TestPramSign(t *testing.T) {
	sign := PramSign([]string{"b3a439d2f3ff5e663152a0efa8283e4f",
		"100096", "BtgPayPlugin",
		"1677833843", "fea5657f-14f0-42e1-890c-93fb08e9afee"})
	fmt.Println(sign)
}
