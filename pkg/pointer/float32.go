package pointer


// ToFloat32Pointer 将float32类型的变量转换为对应的*float32指针类型
func ToFloat32Pointer(v float32) *float32 {
	return &v
}

// ToFloat32PointerOrNilIfZero 将float32类型的变量转换为对应的*float32指针类型，如果变量的值为0的话则返回nil指针
func ToFloat32PointerOrNilIfZero(v float32) *float32 {
	if v < 0.0000001 {
		return nil
	}
	return &v
}

// FromFloat32Pointer 获取*float32类型的指针的实际值，如果指针为nil的话则返回0
func FromFloat32Pointer(p *float32) float32 {
	return FromFloat32PointerOrDefaultIfNil(p, 0)
}

// FromFloat32PointerOrDefaultIfNil 获取*float32类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromFloat32PointerOrDefaultIfNil(v *float32, defaultValue float32) float32 {
	if v == nil {
		return defaultValue
	}
	return *v
}
