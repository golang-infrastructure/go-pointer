package pointer

// ToIntPointer 将int类型的变量转换为对应的*int指针类型
func ToIntPointer(v int) *int {
	return &v
}

// ToIntPointerOrNilIfZero 将int类型的变量转换为对应的*int指针类型，如果变量的值为0的话则返回nil指针
func ToIntPointerOrNilIfZero(v int) *int {
	if v == 0 {
		return nil
	}
	return &v
}

// FromIntPointer 获取*int类型的指针的实际值，如果指针为nil的话则返回0
func FromIntPointer(p *int) int {
	return FromIntPointerOrDefaultIfNil(p, 0)
}

// FromIntPointerOrDefaultIfNil 获取*int类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromIntPointerOrDefaultIfNil(v *int, defaultValue int) int {
	if v == nil {
		return defaultValue
	}
	return *v
}
