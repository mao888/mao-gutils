package concurrent

import (
	"fmt"
	"testing"
)

func TestConcurrentExecuteBetter(t *testing.T) {
	var (
		uids    = []int64{1, 2, 3, 4, 5}
		result  ConcurrentResultBetter[int64]
		mapping map[int64]int64
	)

	if result = ConcurrentExecuteBetter(getUID, uids, 5); result.ExistError {
		t.Errorf("ConcurrentExecuteBetter failed, error count: %d", result.ErrorCount)
		return
	}

	mapping = make(map[int64]int64, len(uids))
	for index, item := range result.Results {
		mapping[uids[index]] = item.Value
	}
	fmt.Printf("mapping: %v", mapping)
}

func getUID(uid int64) (int64, error) {
	return uid, nil
}
