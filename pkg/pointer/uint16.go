package pointer


// ToUInt16Pointer 将uint16类型的变量转换为对应的*uint16指针类型
func ToUInt16Pointer(v uint16) *uint16 {
	return &v
}

// ToUInt16PointerOrNilIfZero 将uint16类型的变量转换为对应的*uint16指针类型，如果变量的值为0的话则返回nil指针
func ToUInt16PointerOrNilIfZero(v uint16) *uint16 {
	if v == 0 {
		return nil
	}
	return &v
}

// FromUInt16Pointer 获取*uint16类型的指针的实际值，如果指针为nil的话则返回0
func FromUInt16Pointer(p *uint16) uint16 {
	return FromUInt16PointerOrDefaultIfNil(p, 0)
}

// FromUInt16PointerOrDefaultIfNil 获取*uint16类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromUInt16PointerOrDefaultIfNil(v *uint16, defaultValue uint16) uint16 {
	if v == nil {
		return defaultValue
	}
	return *v
}