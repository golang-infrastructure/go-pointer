# Go Pointer

在golang中基本类型因为没有包装类型，所以很多库都倾向于采用基本类型变量的指针来区分是没有传递还是传递了零值。

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

但是有时候这个值就是一个字面值常量传进去的，这个时候如果要获取指针类型的话就有点麻烦，上面这个场景只是举了一个例子，这个模块就是用来解决类似的问题的。

目前支持的数据类型：

```text
bool
byte
complex64
complex128
float32
float64
int
int8
int16
int32
int64
rune
string
time.Duration
time.Time
uint
uint8
uint16
uint32
uint64
```

# string

# int

# float32

# float64

# bool

```go

```

# byte

#  



