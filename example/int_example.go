package example

import (
	"fmt"
	"github.com/CC11001100/go-pointer/pkg/pointer"
	"reflect"
)

// IntExample int与*int类型互相转换的示例代码
func IntExample() {

	var intVar int
	var intPointer *int

	// 将int变量转为*int类型的指针
	intVar = 10086
	intPointer = pointer.ToIntPointer(intVar)
	fmt.Println(reflect.TypeOf(intPointer)) // output: *int
	fmt.Println(*intPointer)                // output: 10086

	// 可以在int转换为*int的时候把0值转为nil指针
	intVar = 0
	intPointer = pointer.ToIntPointerOrNilIfZero(intVar)
	fmt.Println(intPointer) // output: <nil>

	// 读取*int指针所指向的地址的值
	intPointer = pointer.ToIntPointer(10010)
	intVar = pointer.FromIntPointer(intPointer)
	fmt.Println(intVar) // output: 10010

	// 如果*int为nil，则返回0
	intPointer = nil
	intVar = pointer.FromIntPointer(intPointer)
	fmt.Println(intVar) // output: 0

	// 可以自定义为nil的时候的默认值
	intVar = pointer.FromIntPointerOrDefaultIfNil(intPointer, 100)
	fmt.Println(intVar) // output: 100

}
