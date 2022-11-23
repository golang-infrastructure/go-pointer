# Go Pointer

[中文文档](./README.md) [English Document](./README_en.md)

# 一、引入依赖

```text
go get -u github.com/golang-infrastructure/go-pointer
```

# 二、解决了什么问题

在golang中基本类型因为没有包装类型，这就导致基本类型无法区分出nil和零值，于是所以很多库都倾向于采用基本类型变量的指针来区分是没有传递还是传递了零值。

一个具体的例子，当不使用指针的时候，在执行任务的时候有个配置项：

```go
package main

type Config struct {
	Foo int
}

```

当Foo的值为0的时候我们不知道是传递了零值还是没有传递值，因为有些地方是要比较细的区分这些情况的，这个时候一些库就倾向于采用指针类型：

```go
package main

type Config struct {
	Foo *int
}

```

但是有时候这个值就是一个字面值常量传进去的，比如查询数据库时的分页大小等，这个时候如果要获取指针类型的话就有点麻烦，比如：

```go
package main 

func main() {
    
    foo := 10 
    config := &Config{
        Foo: &foo, 
    }
    callSomeFunction(config)
    
}
```

如果使用这个库的话：

```go
package main 

func main() {
    
    config := &Config{
        Foo: pointer.ToPointer(10) 
    }
    callSomeFunction(config)
    
}
```

Diff:

```diff
package main 

func main() {

-	foo := 10 
    config := &Config{
-    	Foo: &foo, 
+       Foo: pointer.ToPointer(10) 
    }
    callSomeFunction(config)
    
}
```

上面这个场景只是举了一个例子，这个模块就是用来解决类似的问题的。

# 三、Example Code

目前已经支持泛型：

```go
package main

import (
	"fmt"
	pointer "github.com/golang-infrastructure/go-reflect-utils"
)

func main() {

	// 返回false的指针
	falsePointer := pointer.FalsePointer()
	fmt.Println(fmt.Sprintf("%T %v", falsePointer, *falsePointer)) // Output: *bool false

	// 返回true的指针
	truePointer := pointer.TruePointer()
	fmt.Println(fmt.Sprintf("%T %v", truePointer, *truePointer)) // Output: *bool true

	// 返回对应类型的指针
	v1 := 1
	toPointer := pointer.ToPointer(v1)
	fmt.Println(fmt.Sprintf("%T %v", toPointer, *toPointer)) // Output: *int 1
	// 返回对应类型的指针，但是会检查值，如果值为对应类型的零值的话则返回nil
	v1 = 0
	orNil := pointer.ToPointerOrNil(v1)
	fmt.Println(orNil) // Output: nil

	// 从指针读取值 ，如果为nil指针的话则返回对应类型的零值
	v2 := 1
	v3 := &v2
	fromPointer := pointer.FromPointer(v3)
	fmt.Println(fromPointer) // Output: 1
	// 从指针读取值，如果是nil指针的话则返回给定的默认值
	v2 = 0
	orDefault := pointer.FromPointerOrDefault(v3, 1)
	fmt.Println(orDefault) // Output: 0
}
```



