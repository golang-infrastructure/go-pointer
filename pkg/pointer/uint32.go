package pointer


// ToUInt32Pointer 将uint32类型的变量转换为对应的*uint32指针类型
func ToUInt32Pointer(v uint32) *uint32 {
	return &v
}

// ToUInt32PointerOrNilIfZero 将uint32类型的变量转换为对应的*uint32指针类型，如果变量的值为0的话则返回nil指针
func ToUInt32PointerOrNilIfZero(v uint32) *uint32 {
	if v == 0 {
		return nil
	}
	return &v
}

// FromUInt32Pointer 获取*uint32类型的指针的实际值，如果指针为nil的话则返回0
func FromUInt32Pointer(p *uint32) uint32 {
	return FromUInt32PointerOrDefaultIfNil(p, 0)
}

// FromUInt32PointerOrDefaultIfNil 获取*uint32类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromUInt32PointerOrDefaultIfNil(v *uint32, defaultValue uint32) uint32 {
	if v == nil {
		return defaultValue
	}
	return *v
}