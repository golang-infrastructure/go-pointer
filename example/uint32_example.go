package example

import (
	"fmt"
	"github.com/CC11001100/go-pointer/pkg/pointer"
	"reflect"
)

// UInt32Example uint32与*uint32类型互相转换的示例代码
func UInt32Example() {

	var uint32Var uint32
	var uint32Pointer *uint32

	// 将uint32变量转为*uint32类型的指针
	uint32Var = 10086
	uint32Pointer = pointer.ToUInt32Pointer(uint32Var)
	fmt.Println(reflect.TypeOf(uint32Pointer)) // output: *uint32
	fmt.Println(*uint32Pointer)                // output: 10086

	// 可以在uint32转换为*uint32的时候把0值转为nil指针
	uint32Var = 0
	uint32Pointer = pointer.ToUInt32PointerOrNilIfZero(uint32Var)
	fmt.Println(uint32Pointer) // output: <nil>

	// 读取*uint32指针所指向的地址的值
	uint32Pointer = pointer.ToUInt32Pointer(10010)
	uint32Var = pointer.FromUInt32Pointer(uint32Pointer)
	fmt.Println(uint32Var) // output: 10010

	// 如果*uint32为nil，则返回0
	uint32Pointer = nil
	uint32Var = pointer.FromUInt32Pointer(uint32Pointer)
	fmt.Println(uint32Var) // output: 0

	// 可以自定义为nil的时候的默认值
	uint32Var = pointer.FromUInt32PointerOrDefaultIfNil(uint32Pointer, 100)
	fmt.Println(uint32Var) // output: 100

}
