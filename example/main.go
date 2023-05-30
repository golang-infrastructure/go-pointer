package main

import (
	"fmt"
	"github.com/golang-infrastructure/go-pointer"
)

func main() {

	// 返回false的指针
	falsePointer := pointer.FalsePointer()
	fmt.Println(fmt.Sprintf("%T %v", falsePointer, *falsePointer)) // Output: *bool false

	// 返回true的指针
	truePointer := pointer.TruePointer()
	fmt.Println(fmt.Sprintf("%T %v", truePointer, *truePointer)) // Output: *bool true

	// 返回对应类型的指针
	v1 := 1
	toPointer := pointer.ToPointer(v1)
	fmt.Println(fmt.Sprintf("%T %v", toPointer, *toPointer)) // Output: *int 1
	// 返回对应类型的指针，但是会检查值，如果值为对应类型的零值的话则返回nil
	v1 = 0
	orNil := pointer.ToPointerOrNil(v1)
	fmt.Println(orNil) // Output: nil

	// 从指针读取值 ，如果为nil指针的话则返回对应类型的零值
	v2 := 1
	v3 := &v2
	fromPointer := pointer.FromPointer(v3)
	fmt.Println(fromPointer) // Output: 1
	// 从指针读取值，如果是nil指针的话则返回给定的默认值
	v2 = 0
	orDefault := pointer.FromPointerOrDefault(v3, 1)
	fmt.Println(orDefault) // Output: 0

	// 返回当前时间的指针
	nowPointer := pointer.Now()
	fmt.Println(nowPointer) // Output: 2023-05-30 11:46:20.3695476 +0800 CST m=+0.003922101
}
