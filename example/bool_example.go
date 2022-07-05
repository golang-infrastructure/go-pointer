package example

import (
	"fmt"
	"go-pointer/pkg/pointer"
)

func BoolExample() {

	// 获取一个布尔指针，其值为true
	boolPointer := pointer.TruePointer()
	fmt.Println(*boolPointer)

	// 获取一个布尔指针，其值为false
	boolPointer = pointer.FalsePointer()
	fmt.Println(*boolPointer)



}
