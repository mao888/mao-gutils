package os

import (
	"fmt"
	"testing"
)

// 使用map替换template模版中 $ 符号后的字符串
func TestExpandByMap(t *testing.T) {
	template := "chartid=$user_id & prop=$prop"

	// 声明map
	var mapStr map[string]string
	//使用make函数创建一个非nil的map，nil map不能赋值
	mapStr = make(map[string]string)
	//给已声明的map赋值
	mapStr["user_id"] = "123"
	mapStr["prop"] = "huchao"

	//res := os.Expand(template, func(s string) string {
	//	return mapStr[s]
	//})

	res := ExpandByMap(template, mapStr)
	fmt.Println(res)
}

// 自定义函数替换字符串模版中 $ 符号后的字符串
func TestExpandByFun(t *testing.T) {
	str := "My name is $Name and live in $City."

	//Format the specified string using os.Expand
	strResult := ExpandByFun(str, GetVal)

	fmt.Println(strResult)
}

func GetVal(val string) string {
	switch val {
	case "Name":
		return "Morrish"
	case "City":
		return "New York"
	default:
		return "<Empty>"
	}
}

// 替换template模版中 $ 符号后的字符串
func TestGetComposedTemplateListExpandByMap(t *testing.T) {
	template := "chartid=$user_id & prop=$prop"

	// 声明map
	var mapStr map[string]string
	//使用make函数创建一个非nil的map，nil map不能赋值
	mapStr = make(map[string]string)
	//给已声明的map赋值
	mapStr["user_id"] = "1,2,3"
	mapStr["prop"] = "huchao"

	strArr := GetComposedTemplateListExpandByMap(template, true, mapStr)
	fmt.Println("strArr", strArr) // [chartid=1 & prop=huchao chartid=2 & prop=huchao chartid=3 & prop=huchao]
}
