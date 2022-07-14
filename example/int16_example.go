package example

import (
	"fmt"
	"github.com/CC11001100/go-pointer/pkg/pointer"
	"reflect"
)

// Int16Example int16与*int16类型互相转换的示例代码
func Int16Example() {

	var int16Var int16
	var int16Pointer *int16

	// 将int16变量转为*int16类型的指针
	int16Var = 10086
	int16Pointer = pointer.ToInt16Pointer(int16Var)
	fmt.Println(reflect.TypeOf(int16Pointer)) // output: *int16
	fmt.Println(*int16Pointer)                // output: 10086

	// 可以在int16转换为*int16的时候把0值转为nil指针
	int16Var = 0
	int16Pointer = pointer.ToInt16PointerOrNilIfZero(int16Var)
	fmt.Println(int16Pointer) // output: <nil>

	// 读取*int16指针所指向的地址的值
	int16Pointer = pointer.ToInt16Pointer(10010)
	int16Var = pointer.FromInt16Pointer(int16Pointer)
	fmt.Println(int16Var) // output: 10010

	// 如果*int16为nil，则返回0
	int16Pointer = nil
	int16Var = pointer.FromInt16Pointer(int16Pointer)
	fmt.Println(int16Var) // output: 0

	// 可以自定义为nil的时候的默认值
	int16Var = pointer.FromInt16PointerOrDefaultIfNil(int16Pointer, 100)
	fmt.Println(int16Var) // output: 100

}
