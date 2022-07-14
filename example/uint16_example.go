package example

import (
	"fmt"
	"github.com/CC11001100/go-pointer/pkg/pointer"
	"reflect"
)

// UInt16Example uint16与*uint16类型互相转换的示例代码
func UInt16Example() {

	var uint16Var uint16
	var uint16Pointer *uint16

	// 将uint16变量转为*uint16类型的指针
	uint16Var = 10086
	uint16Pointer = pointer.ToUInt16Pointer(uint16Var)
	fmt.Println(reflect.TypeOf(uint16Pointer)) // output: *uint16
	fmt.Println(*uint16Pointer)                // output: 10086

	// 可以在uint16转换为*uint16的时候把0值转为nil指针
	uint16Var = 0
	uint16Pointer = pointer.ToUInt16PointerOrNilIfZero(uint16Var)
	fmt.Println(uint16Pointer) // output: <nil>

	// 读取*uint16指针所指向的地址的值
	uint16Pointer = pointer.ToUInt16Pointer(10010)
	uint16Var = pointer.FromUInt16Pointer(uint16Pointer)
	fmt.Println(uint16Var) // output: 10010

	// 如果*uint16为nil，则返回0
	uint16Pointer = nil
	uint16Var = pointer.FromUInt16Pointer(uint16Pointer)
	fmt.Println(uint16Var) // output: 0

	// 可以自定义为nil的时候的默认值
	uint16Var = pointer.FromUInt16PointerOrDefaultIfNil(uint16Pointer, 100)
	fmt.Println(uint16Var) // output: 100

}
