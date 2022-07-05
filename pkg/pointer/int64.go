package pointer

// ToInt64Pointer 将int64类型的变量转换为对应的*int64指针类型
func ToInt64Pointer(v int64) *int64 {
	return &v
}

// ToInt64PointerOrNilIfZero 将int64类型的变量转换为对应的*int64指针类型，如果变量的值为0的话则返回nil指针
func ToInt64PointerOrNilIfZero(v int64) *int64 {
	if v == 0 {
		return nil
	}
	return &v
}

// FromInt64Pointer 获取*int64类型的指针的实际值，如果指针为nil的话则返回0
func FromInt64Pointer(p *int64) int64 {
	return FromInt64PointerOrDefaultIfNil(p, 0)
}

// FromInt64PointerOrDefaultIfNil 获取*int64类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromInt64PointerOrDefaultIfNil(v *int64, defaultValue int64) int64 {
	if v == nil {
		return defaultValue
	}
	return *v
}
