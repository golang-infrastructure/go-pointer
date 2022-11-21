package pointer

import reflectutils "github.com/golang-infrastructure/go-reflect-utils"

// TruePointer 返回一个布尔指针，其值为true
func TruePointer() *bool {
	b := true
	return &b
}

// FalsePointer 返回一个布尔指针，其值为false
func FalsePointer() *bool {
	b := false
	return &b
}

// ToPointer 将布尔变量转换为布尔指针
func ToPointer[T any](value T) *T {
	return &value
}

// ToPointerOrNil 转换为对应类型的指针，如果变量的值为对应类型的零值的话则转换为nil指针，如果是struct，则每个字段都为零值
func ToPointerOrNil[T any](value T) *T {
	if reflectutils.IsZero(value) {
		return nil
	} else {
		return &value
	}
}

// FromPointer 获取指针实际指向的值，如果指针为nil的话则返回对应类型的零值
func FromPointer[T any](valuePointer *T) T {
	var zero T
	return FromPointerOrDefault(valuePointer, zero)
}

// FromPointerOrDefault 获取指针实际指向的值，如果指针为nil的话则返回给定的defaultValue
func FromPointerOrDefault[T any](valuePointer *T, defaultValue T) T {
	if reflectutils.IsNil(valuePointer) {
		return defaultValue
	} else {
		return *valuePointer
	}
}
