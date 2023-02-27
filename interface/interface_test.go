// Golang program to illustrate the
// concept of type switches
package main

import (
	"fmt"
	"testing"
)

func TestJudgeType(t *testing.T) {
	// an interface that has
	// a string value
	var (
		value  interface{} = "GeeksforGeeks" // string
		value2 interface{} = 1               // int
		value3 interface{} = 1.11            // float64
		value4 interface{} = true            // bool
		value5 interface{} = nil             // nil
		value6 interface{} = ""              // string
	)

	fmt.Printf("type of %s: %s \n ", value, JudgeType(value))
	fmt.Printf("type of %d: %s \n ", value2, JudgeType(value2))
	fmt.Printf("type of %f: %s \n ", value3, JudgeType(value3))
	fmt.Printf("type of %v: %s \n ", value4, JudgeType(value4))
	fmt.Printf("type of %v: %s \n ", value5, JudgeType(value5))
	fmt.Printf("type of %v: %s \n ", value6, JudgeType(value6))
}

func TestJudgeTypeByReflect(t *testing.T) {
	// an interface that has
	// a string value
	var (
		value  interface{} = "GeeksforGeeks" // string
		value2 interface{} = 1               // int
		value3 interface{} = 1.11            // float64
		value4 interface{} = true            // bool
		//value5 interface{} = nil             // nil
		value6 interface{} = "" // string
	)

	fmt.Printf("type of %s: %s \n ", value, JudgeTypeByReflect(value))
	fmt.Printf("type of %d: %s \n ", value2, JudgeTypeByReflect(value2))
	fmt.Printf("type of %f: %s \n ", value3, JudgeTypeByReflect(value3))
	fmt.Printf("type of %v: %s \n ", value4, JudgeTypeByReflect(value4))
	//fmt.Printf("type of %v: %s \n ", value5, JudgeTypeByReflect(value5))
	fmt.Printf("type of %v: %s \n ", value6, JudgeTypeByReflect(value6))
}
