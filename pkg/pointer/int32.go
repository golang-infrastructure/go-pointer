package pointer

// ToInt32Pointer 将int32类型的变量转换为对应的*int32指针类型
func ToInt32Pointer(v int32) *int32 {
	return &v
}

// ToInt32PointerOrNilIfZero 将int32类型的变量转换为对应的*int32指针类型，如果变量的值为0的话则返回nil指针
func ToInt32PointerOrNilIfZero(v int32) *int32 {
	if v == 0 {
		return nil
	}
	return &v
}

// FromInt32Pointer 获取*int32类型的指针的实际值，如果指针为nil的话则返回0
func FromInt32Pointer(p *int32) int32 {
	return FromInt32PointerOrDefaultIfNil(p, 0)
}

// FromInt32PointerOrDefaultIfNil 获取*int32类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromInt32PointerOrDefaultIfNil(v *int32, defaultValue int32) int32 {
	if v == nil {
		return defaultValue
	}
	return *v
}
