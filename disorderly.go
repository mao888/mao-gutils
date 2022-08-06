package utils

import (
	"math/rand"
	"reflect"
	"time"
)

// RandSlice 打乱顺序
func RandSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	// 通过反射获取是否属于切片类型
	if rv.Type().Kind() != reflect.Slice {
		return
	}

	// 获取长度，如果不足两个没必要继续执行
	length := rv.Len()
	if length < 2 {
		return
	}
	// 通过Swapper用于交换切片中的元素
	swap := reflect.Swapper(slice)
	rand.Seed(time.Now().Unix()) // 用时间戳来生成随机种子保证每次随机种子不同
	for i := length - 1; i >= 0; i-- {
		j := rand.Intn(length) // 在切片长度内随机生成一个数字
		swap(i, j)             // 传入下标进行交换
	}
	return
}
