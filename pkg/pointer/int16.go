package pointer

// ToInt16Pointer 将int16类型的变量转换为对应的*int16指针类型
func ToInt16Pointer(v int16) *int16 {
	return &v
}

// ToInt16PointerOrNilIfZero 将int16类型的变量转换为对应的*int16指针类型，如果变量的值为0的话则返回nil指针
func ToInt16PointerOrNilIfZero(v int16) *int16 {
	if v == 0 {
		return nil
	}
	return &v
}

// FromInt16Pointer 获取*int16类型的指针的实际值，如果指针为nil的话则返回0
func FromInt16Pointer(p *int16) int16 {
	return FromInt16PointerOrDefaultIfNil(p, 0)
}

// FromInt16PointerOrDefaultIfNil 获取*int16类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromInt16PointerOrDefaultIfNil(v *int16, defaultValue int16) int16 {
	if v == nil {
		return defaultValue
	}
	return *v
}
