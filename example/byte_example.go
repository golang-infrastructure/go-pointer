package example

//func ByteExample() {
//
//	// 获取一个布尔指针，其值为true
//	boolPointer := pointer.Byte()
//	fmt.Println(reflect.TypeOf(boolPointer)) // output: *bool
//	fmt.Println(*boolPointer)                // output: true
//
//	// 获取一个布尔指针，其值为false
//	boolPointer = pointer.FalsePointer()
//	fmt.Println(reflect.TypeOf(boolPointer)) // output: *bool
//	fmt.Println(*boolPointer)                // output: false
//
//	// 将布尔变量转换为布尔指针
//	boolVar := true
//	boolPointer = pointer.ToBoolPointer(boolVar)
//	fmt.Println(reflect.TypeOf(boolPointer)) // output: *bool
//	fmt.Println(*boolPointer)                // output: true
//	boolVar = false
//	boolPointer = pointer.ToBoolPointer(boolVar)
//	fmt.Println(reflect.TypeOf(boolPointer)) // output: *bool
//	fmt.Println(*boolPointer)                // output: false
//
//	// 将布尔变量转换为布尔指针，如果布尔的值为nil的话则返回nil指针
//	boolVar = true
//	boolPointer = pointer.ToBoolPointerOrNilIfFalse(boolVar)
//	fmt.Println(reflect.TypeOf(boolPointer)) // output: *bool
//	fmt.Println(*boolPointer)                // output: true
//	boolVar = false
//	boolPointer = pointer.ToBoolPointerOrNilIfFalse(boolVar)
//	fmt.Println(reflect.TypeOf(boolPointer)) // output: *bool
//	fmt.Println(boolPointer)                 // output: <nil>
//
//}
