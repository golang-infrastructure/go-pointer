package example

import (
	"fmt"
	"github.com/CC11001100/go-pointer/pkg/pointer"
	"reflect"
	"time"
)

func TimeDurationExample() {

	var durationVar time.Duration
	var durationPointer *time.Duration

	// 把time.Duration转为*time.Duration
	durationVar = time.Duration(100)
	durationPointer = pointer.ToDurationPointer(durationVar)
	fmt.Println(reflect.TypeOf(durationPointer)) // output:
	fmt.Println(*durationPointer)                // output:

	// 

}
