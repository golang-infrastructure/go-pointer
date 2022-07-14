package example

import (
	"fmt"
	"github.com/CC11001100/go-pointer/pkg/pointer"
	"reflect"
)

func BoolExample() {

	// 获取一个布尔指针，其值为true
	boolPointer := pointer.TruePointer()
	fmt.Println(reflect.TypeOf(boolPointer)) // output: *bool
	fmt.Println(*boolPointer)                // output: true

	// 获取一个布尔指针，其值为false
	boolPointer = pointer.FalsePointer()
	fmt.Println(reflect.TypeOf(boolPointer)) // output: *bool
	fmt.Println(*boolPointer)                // output: false

	// 将布尔变量转换为布尔指针
	boolVar := true
	boolPointer = pointer.ToBoolPointer(boolVar)
	fmt.Println(reflect.TypeOf(boolPointer)) // output: *bool
	fmt.Println(*boolPointer)                // output: true
	boolVar = false
	boolPointer = pointer.ToBoolPointer(boolVar)
	fmt.Println(reflect.TypeOf(boolPointer)) // output: *bool
	fmt.Println(*boolPointer)                // output: false

	// 将布尔变量转换为布尔指针，如果布尔的值为nil的话则返回nil指针
	boolVar = true
	boolPointer = pointer.ToBoolPointerOrNilIfFalse(boolVar)
	fmt.Println(reflect.TypeOf(boolPointer)) // output: *bool
	fmt.Println(*boolPointer)                // output: true
	boolVar = false
	boolPointer = pointer.ToBoolPointerOrNilIfFalse(boolVar)
	fmt.Println(reflect.TypeOf(boolPointer)) // output: *bool
	fmt.Println(boolPointer)                 // output: <nil>

	// 读取布尔指针变量的值，如果值为nil的话则返回false
	boolVar = pointer.FromBoolPointer(pointer.FalsePointer())
	fmt.Println(boolVar)

	// 读取布尔变量的值，如果布尔指针为nil的话则返回给定的默认值，这个默认值可以是true或者false
	boolPointer = nil
	boolVar = pointer.FromBoolPointerOrDefault(boolPointer, true)
	fmt.Println(boolVar) // output: true

}
