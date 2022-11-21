# Go Pointer

[中文文档](./README.md) [English Document](./README_en.md)

# 1. Use go get install it

```text
go get -u github.com/golang-infrastructure/go-pointer
```

# 2. What problem was solved

In golang, the primitive type has no wrapping type, so the primitive type can't distinguish between nil and zero,
so many libraries tend to use Pointers to primitive variables to distinguish between zero and not passed.

For a specific example, when no Pointers are used, there is a configuration item when executing a task:

```go
package main

type Config struct {
	Foo int
}

```

When Foo is 0, we do not know whether we have passed zero or no value, because some libraries prefer to use pointer
types in this case:

```go
package main

type Config struct {
	Foo *int
}

```

However, sometimes this value is passed in as a literal constant, such as the paging size when querying a database. In
this case,
getting a pointer type can be a bit of a hassle. The above scenario is just an example of a problem that this module is
designed to solve.

# 3. Example Code

Generics are already supported:

```go
package main

import (
	"fmt"
	pointer "github.com/golang-infrastructure/go-reflect-utils"
)

func main() {

	// Returns a pointer to false
	falsePointer := pointer.FalsePointer()
	fmt.Println(fmt.Sprintf("%T %v", falsePointer, *falsePointer)) // Output: *bool false

	// Returns a pointer to true
	truePointer := pointer.TruePointer()
	fmt.Println(fmt.Sprintf("%T %v", truePointer, *truePointer)) // Output: *bool true

	// Returns a pointer to the corresponding type
	v1 := 1
	toPointer := pointer.ToPointer(v1)
	fmt.Println(fmt.Sprintf("%T %v", toPointer, *toPointer)) // Output: *int 1
	// Returns a pointer to the corresponding type, but checks the value and returns nil if the value is zero of the corresponding type
	v1 = 0
	orNil := pointer.ToPointerOrNil(v1)
	fmt.Println(orNil) // Output: nil

	// Reads a value from a pointer, and returns a zero value of the corresponding type if it is a nil pointer
	v2 := 1
	v3 := &v2
	fromPointer := pointer.FromPointer(v3)
	fmt.Println(fromPointer) // Output: 1
	// Reads the value from a pointer, and returns the given default value if it is a nil pointer
	v2 = 0
	orDefault := pointer.FromPointerOrDefault(v3, 1)
	fmt.Println(orDefault) // Output: 0
}
```




