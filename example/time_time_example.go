package example

import (
	"fmt"
	"github.com/CC11001100/go-pointer/pkg/pointer"
	"reflect"
	"time"
)

func TimeTimeExample() {

	var timeVar time.Time
	var timePointer *time.Time

	// 获取当前时间的*time.Time指针
	timePointer = pointer.NowTimePointer()
	fmt.Println(reflect.TypeOf(timePointer)) // output: *time.Time
	fmt.Println(*timePointer)                // output: 2022-07-06 01:52:51.6130516 +0800 CST m=+0.008121201

	// 将time.Time变量转换为*time.Time指针
	timeVar = time.Now()
	timePointer = pointer.ToTimePointer(timeVar)
	fmt.Println(reflect.TypeOf(timePointer)) // output: *time.Time
	fmt.Println(*timePointer)                // output: 2022-07-06 01:52:51.6410163 +0800 CST m=+0.036085901

	// 当把time.Time转为*time.Time的时候，如果传入的time.Time是零值，则返回nil
	timeVar = time.Time{}
	timePointer = pointer.ToTimePointer(timeVar)
	fmt.Println(timePointer) // output: <nil>

	// 读取*time.Time的值
	timePointer = pointer.NowTimePointer()
	timeVar = pointer.FromTimePointer(timePointer)
	fmt.Println(timeVar) // output: 2022-07-06 01:52:51.6410163 +0800 CST m=+0.036085901

	// 要读取*time.Time的值的时候，指针为nil会返回time.Time的零值
	timePointer = nil
	timeVar = pointer.FromTimePointer(timePointer)
	fmt.Println(timeVar)          // output: 0001-01-01 00:00:00 +0000 UTC
	fmt.Println(timeVar.IsZero()) // output: true

	// 可以制定当*time.Time为nil的时候的默认值
	timePointer = nil
	timeVar = pointer.FromTimePointerOrDefault(timePointer, time.Now())
	fmt.Println(timeVar) // output: 2022-07-06 01:52:51.6410163 +0800 CST m=+0.036085901

}
