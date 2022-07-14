package example

import (
	"fmt"
	"github.com/CC11001100/go-pointer/pkg/pointer"
	"reflect"
)

// UInt64Example uint64与*uint64类型互相转换的示例代码
func UInt64Example() {

	var uint64Var uint64
	var uint64Pointer *uint64

	// 将uint64变量转为*uint64类型的指针
	uint64Var = 10086
	uint64Pointer = pointer.ToUInt64Pointer(uint64Var)
	fmt.Println(reflect.TypeOf(uint64Pointer)) // output: *uint64
	fmt.Println(*uint64Pointer)                // output: 10086

	// 可以在uint64转换为*uint64的时候把0值转为nil指针
	uint64Var = 0
	uint64Pointer = pointer.ToUInt64PointerOrNilIfZero(uint64Var)
	fmt.Println(uint64Pointer) // output: <nil>

	// 读取*uint64指针所指向的地址的值
	uint64Pointer = pointer.ToUInt64Pointer(10010)
	uint64Var = pointer.FromUInt64Pointer(uint64Pointer)
	fmt.Println(uint64Var) // output: 10010

	// 如果*uint64为nil，则返回0
	uint64Pointer = nil
	uint64Var = pointer.FromUInt64Pointer(uint64Pointer)
	fmt.Println(uint64Var) // output: 0

	// 可以自定义为nil的时候的默认值
	uint64Var = pointer.FromUInt64PointerOrDefaultIfNil(uint64Pointer, 100)
	fmt.Println(uint64Var) // output: 100

}
