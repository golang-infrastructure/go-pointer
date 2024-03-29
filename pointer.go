package pointer

import (
	reflectutils "github.com/golang-infrastructure/go-reflect-utils"
	"sync/atomic"
	"time"
)

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

// Now 返回当前时间的指针
func Now() *time.Time {
	return ToPointer(time.Now())
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

func AtomicTruePointer() *atomic.Bool {
	b := &atomic.Bool{}
	b.Store(true)
	return b
}

func AtomicFalsePointer() *atomic.Bool {
	b := &atomic.Bool{}
	b.Store(false)
	return b
}

func AtomicUInt64Pointer(value uint64) *atomic.Uint64 {
	i := &atomic.Uint64{}
	i.Store(value)
	return i
}

func AtomicUInt32Pointer(value uint32) *atomic.Uint32 {
	i := &atomic.Uint32{}
	i.Store(value)
	return i
}

func AtomicInt64Pointer(value int64) *atomic.Int64 {
	i := &atomic.Int64{}
	i.Store(value)
	return i
}

func AtomicInt32Pointer(value int32) *atomic.Int32 {
	i := &atomic.Int32{}
	i.Store(value)
	return i
}
