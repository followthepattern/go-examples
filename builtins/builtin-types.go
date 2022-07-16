package builtins

import (
	"fmt"
	"reflect"
)

func BuiltInExamples() {
	// you can declare a type and give value later
	var value1 string
	value1 = "hello"
	fmt.Println(reflect.TypeOf(value1))

	value2 := "hello" == "helo" // bool
	fmt.Println(reflect.TypeOf(value2))

	value3 := 1 // int
	fmt.Println(reflect.TypeOf(value3))

	// different int types must be explicitly declared
	var value4 int8 = 1
	fmt.Println(reflect.TypeOf(value4))

	var value5 int64 = 1
	fmt.Println(reflect.TypeOf(value5))

	var value6 uint = 1
	fmt.Println(reflect.TypeOf(value6))

	var value7 uint16 = 23
	fmt.Println(reflect.TypeOf(value7))

	value8 := 8.7 // float64
	fmt.Println(reflect.TypeOf(value8))

	var value9 float32 = 382.394
	fmt.Println(reflect.TypeOf(value9))
}
