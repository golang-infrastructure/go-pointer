package example

import (
	"fmt"
	"github.com/CC11001100/go-pointer/pkg/pointer"
	"reflect"
)

// Int8Example int8与*int8类型互相转换的示例代码
func Int8Example() {

	var int8Var int8
	var int8Pointer *int8

	// 将int8变量转为*int8类型的指针
	int8Var = 64
	int8Pointer = pointer.ToInt8Pointer(int8Var)
	fmt.Println(reflect.TypeOf(int8Pointer)) // output: *int8
	fmt.Println(*int8Pointer)                // output: 64

	// 可以在int8转换为*int8的时候把0值转为nil指针
	int8Var = 0
	int8Pointer = pointer.ToInt8PointerOrNilIfZero(int8Var)
	fmt.Println(int8Pointer) // output: <nil>

	// 读取*int8指针所指向的地址的值
	int8Pointer = pointer.ToInt8Pointer(32)
	int8Var = pointer.FromInt8Pointer(int8Pointer)
	fmt.Println(int8Var) // output: 32

	// 如果*int8为nil，则返回0
	int8Pointer = nil
	int8Var = pointer.FromInt8Pointer(int8Pointer)
	fmt.Println(int8Var) // output: 0

	// 可以自定义为nil的时候的默认值
	int8Var = pointer.FromInt8PointerOrDefaultIfNil(int8Pointer, 100)
	fmt.Println(int8Var) // output: 100

}
