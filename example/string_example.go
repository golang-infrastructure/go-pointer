package example

import (
	"fmt"
	"github.com/CC11001100/go-pointer/pkg/pointer"
	"reflect"
)

func StringExample() {

	var stringVar string
	var stringPointer *string

	// 把string转为*string
	stringVar = "CC11001100"
	stringPointer = pointer.ToStringPointer(stringVar)
	fmt.Println(reflect.ValueOf(stringPointer)) // output: *string
	fmt.Println(*stringPointer)                 // output: CC11001100

	// 把string转为*string的时候，可以选择如果是空字符串，则将其转为nil
	stringVar = ""
	stringPointer = pointer.ToStringPointerOrNilIfEmpty(stringVar)
	fmt.Println(stringPointer) // output: <nil>

	// 读取*string实际的值，如果为nil的话则返回空字符串
	stringPointer = nil
	stringVar = pointer.FromStringPointer(stringPointer)
	fmt.Println(stringVar) // output: ""

	// 可以自定义在要读取的*string为nil的时候的默认值
	stringPointer = nil
	stringVar = pointer.FromStringPointerOrDefaultIfNil(stringPointer, "CC11001100")
	fmt.Println(stringVar) // output: "CC11001100"

}
