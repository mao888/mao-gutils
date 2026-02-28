package concurrent

import (
	"sync"

	"github.com/mao888/mao-gutils/slice"
)

// ConcurrentFunc 并发执行器
type ConcurrentFunc[T any, R any] func(T) (R, error)

// ConcurrentResult 并发执行器返回数据
// @Date: 2025-09-24 16:02:14
type ConcurrentResult[R any] struct {
	Value R
	Err   error
}

// ConcurrentResultBetter 并发执行器返回数据
type ConcurrentResultBetter[R any] struct {
	ExistError bool                  // 是否存在错误
	ErrorCount int                   // 失败的数量
	Results    []ConcurrentResult[R] // 数据的具体信息
}

// ConcurrentExecute 并发执行
// @params ConcurrentFunc[T, R] 执行器
// @params []T 参数
// @params int 并发数量限制
// @return []ConcurrentResult[R] 返回的数据，这里返回的顺序和参数顺序一致
func ConcurrentExecute[T any, R any](execute ConcurrentFunc[T, R], params []T, limit int) []ConcurrentResult[R] {
	var results []ConcurrentResult[R]
	if len(params) <= 0 {
		return results
	}
	if limit < 1 {
		limit = 1
	}
	results = make([]ConcurrentResult[R], 0, limit)
	batchList := slice.SplitArray(params, limit)
	for _, batch := range batchList {
		if len(batch) > 0 {
			results = append(results, concurrentExecute(execute, batch)...)
		}
	}
	return results
}

// ConcurrentExecuteBetter 并发执行
// @params ConcurrentFunc[T any, R any] 执行器
// @params []T 参数
// @params int 并发数量限制
// @return ConcurrentResultBetter[R] 返回结果
func ConcurrentExecuteBetter[T any, R any](execute ConcurrentFunc[T, R], params []T, limit int) ConcurrentResultBetter[R] {
	var result ConcurrentResultBetter[R]
	if len(params) <= 0 {
		return result
	}
	if limit < 1 {
		limit = 1
	}
	result.Results = make([]ConcurrentResult[R], 0, limit)
	batchList := slice.SplitArray(params, limit)
	for _, batch := range batchList {
		if len(batch) > 0 {
			results, errorCount := concurrentExecuteBetter(execute, batch)
			result.Results = append(result.Results, results...)
			result.ErrorCount += errorCount
		}
	}
	if result.ErrorCount > 0 {
		result.ExistError = true
	}
	return result
}

func concurrentExecute[T any, R any](execute ConcurrentFunc[T, R], params []T) []ConcurrentResult[R] {
	results, _ := concurrentExecuteBetter(execute, params)
	return results
}

func concurrentExecuteBetter[T any, R any](execute ConcurrentFunc[T, R], params []T) ([]ConcurrentResult[R], int) {
	var (
		results    []ConcurrentResult[R]
		mapping    map[int]ConcurrentResult[R]
		errorCount int
		mutex      sync.Mutex
		wait       sync.WaitGroup
	)
	if len(params) <= 0 {
		return results, errorCount
	}

	mapping = make(map[int]ConcurrentResult[R], len(params))
	results = make([]ConcurrentResult[R], 0, len(params))

	wait.Add(len(params))
	for index, param := range params {
		go func(index int, param T) {
			defer wait.Done()
			result, err := execute(param)
			mutex.Lock()
			defer mutex.Unlock()
			mapping[index] = ConcurrentResult[R]{
				Value: result,
				Err:   err,
			}
		}(index, param)
	}
	wait.Wait()
	for index := range params {
		results = append(results, mapping[index])
		if mapping[index].Err != nil {
			errorCount++
		}
	}
	return results, errorCount
}
