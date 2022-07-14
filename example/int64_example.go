package example

import (
	"fmt"
	"github.com/CC11001100/go-pointer/pkg/pointer"
	"reflect"
)

// Int64Example int64与*int64类型互相转换的示例代码
func Int64Example() {

	var int64Var int64
	var int64Pointer *int64

	// 将int64变量转为*int64类型的指针
	int64Var = 10086
	int64Pointer = pointer.ToInt64Pointer(int64Var)
	fmt.Println(reflect.TypeOf(int64Pointer)) // output: *int64
	fmt.Println(*int64Pointer)                // output: 10086

	// 可以在int64转换为*int64的时候把0值转为nil指针
	int64Var = 0
	int64Pointer = pointer.ToInt64PointerOrNilIfZero(int64Var)
	fmt.Println(int64Pointer) // output: <nil>

	// 读取*int64指针所指向的地址的值
	int64Pointer = pointer.ToInt64Pointer(10010)
	int64Var = pointer.FromInt64Pointer(int64Pointer)
	fmt.Println(int64Var) // output: 10010

	// 如果*int64为nil，则返回0
	int64Pointer = nil
	int64Var = pointer.FromInt64Pointer(int64Pointer)
	fmt.Println(int64Var) // output: 0

	// 可以自定义为nil的时候的默认值
	int64Var = pointer.FromInt64PointerOrDefaultIfNil(int64Pointer, 100)
	fmt.Println(int64Var) // output: 100

}
