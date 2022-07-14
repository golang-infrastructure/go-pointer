package example

import (
	"fmt"
	"github.com/CC11001100/go-pointer/pkg/pointer"
	"reflect"
)

// UInt8Example uint8与*uint8类型互相转换的示例代码
func UInt8Example() {

	var uint8Var uint8
	var uint8Pointer *uint8

	// 将uint8变量转为*uint8类型的指针
	uint8Var = 64
	uint8Pointer = pointer.ToUInt8Pointer(uint8Var)
	fmt.Println(reflect.TypeOf(uint8Pointer)) // output: *uint8
	fmt.Println(*uint8Pointer)                // output: 64

	// 可以在uint8转换为*uint8的时候把0值转为nil指针
	uint8Var = 0
	uint8Pointer = pointer.ToUInt8PointerOrNilIfZero(uint8Var)
	fmt.Println(uint8Pointer) // output: <nil>

	// 读取*uint8指针所指向的地址的值
	uint8Pointer = pointer.ToUInt8Pointer(16)
	uint8Var = pointer.FromUInt8Pointer(uint8Pointer)
	fmt.Println(uint8Var) // output: 16

	// 如果*uint8为nil，则返回0
	uint8Pointer = nil
	uint8Var = pointer.FromUInt8Pointer(uint8Pointer)
	fmt.Println(uint8Var) // output: 0

	// 可以自定义为nil的时候的默认值
	uint8Var = pointer.FromUInt8PointerOrDefaultIfNil(uint8Pointer, 100)
	fmt.Println(uint8Var) // output: 100

}
