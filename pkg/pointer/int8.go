package pointer

// ToInt8Pointer 将int8类型的变量转换为对应的*int8指针类型
func ToInt8Pointer(v int8) *int8 {
	return &v
}

// ToInt8PointerOrNilIfZero 将int8类型的变量转换为对应的*int8指针类型，如果变量的值为0的话则返回nil指针
func ToInt8PointerOrNilIfZero(v int8) *int8 {
	if v == 0 {
		return nil
	}
	return &v
}

// FromInt8Pointer 获取*int8类型的指针的实际值，如果指针为nil的话则返回0
func FromInt8Pointer(p *int8) int8 {
	return FromInt8PointerOrDefaultIfNil(p, 0)
}

// FromInt8PointerOrDefaultIfNil 获取*int8类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromInt8PointerOrDefaultIfNil(v *int8, defaultValue int8) int8 {
	if v == nil {
		return defaultValue
	}
	return *v
}
