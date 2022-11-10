package gutil

import (
	"fmt"
	"testing"
)

func TestMapSplitByComma(t *testing.T) {

	// 声明map
	var mapStr map[string]string
	//使用make函数创建一个非nil的map，nil map不能赋值
	mapStr = make(map[string]string)
	//给已声明的map赋值
	mapStr["user_id"] = "1"
	mapStr["prop"] = "huchao"

	mapArr := MapSplitByComma(mapStr)
	fmt.Println("mapArr", mapArr) //  [map[prop:huchao user_id:1] map[prop:huchao user_id:2] map[prop:huchao user_id:3]]
}
