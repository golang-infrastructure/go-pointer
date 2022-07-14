package example

import (
	"fmt"
	"github.com/CC11001100/go-pointer/pkg/pointer"
	"reflect"
)

// Int32Example int32与*int32类型互相转换的示例代码
func Int32Example() {

	var int32Var int32
	var int32Pointer *int32

	// 将int32变量转为*int32类型的指针
	int32Var = 10086
	int32Pointer = pointer.ToInt32Pointer(int32Var)
	fmt.Println(reflect.TypeOf(int32Pointer)) // output: *int32
	fmt.Println(*int32Pointer)                // output: 10086

	// 可以在int32转换为*int32的时候把0值转为nil指针
	int32Var = 0
	int32Pointer = pointer.ToInt32PointerOrNilIfZero(int32Var)
	fmt.Println(int32Pointer) // output: <nil>

	// 读取*int32指针所指向的地址的值
	int32Pointer = pointer.ToInt32Pointer(10010)
	int32Var = pointer.FromInt32Pointer(int32Pointer)
	fmt.Println(int32Var) // output: 10010

	// 如果*int32为nil，则返回0
	int32Pointer = nil
	int32Var = pointer.FromInt32Pointer(int32Pointer)
	fmt.Println(int32Var) // output: 0

	// 可以自定义为nil的时候的默认值
	int32Var = pointer.FromInt32PointerOrDefaultIfNil(int32Pointer, 100)
	fmt.Println(int32Var) // output: 100

}
