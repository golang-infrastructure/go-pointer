package pointer

// ToUInt8Pointer 将uint8类型的变量转换为对应的*uint8指针类型
func ToUInt8Pointer(v uint8) *uint8 {
	return &v
}

// ToUInt8PointerOrNilIfZero 将uint8类型的变量转换为对应的*uint8指针类型，如果变量的值为0的话则返回nil指针
func ToUInt8PointerOrNilIfZero(v uint8) *uint8 {
	if v == 0 {
		return nil
	}
	return &v
}

// FromUInt8Pointer 获取*uint8类型的指针的实际值，如果指针为nil的话则返回0
func FromUInt8Pointer(p *uint8) uint8 {
	return FromUInt8PointerOrDefaultIfNil(p, 0)
}

// FromUInt8PointerOrDefaultIfNil 获取*uint8类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromUInt8PointerOrDefaultIfNil(v *uint8, defaultValue uint8) uint8 {
	if v == nil {
		return defaultValue
	}
	return *v
}
