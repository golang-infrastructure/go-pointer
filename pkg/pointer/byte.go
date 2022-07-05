package pointer

// ToBytePointer 将byte类型的变量转换为对应的*byte指针类型
func ToBytePointer(v byte) *byte {
	return &v
}

// ToBytePointerOrNilIfZero 将byte类型的变量转换为对应的*byte指针类型，如果变量的值为0的话则返回nil指针
func ToBytePointerOrNilIfZero(v byte) *byte {
	if v == 0 {
		return nil
	}
	return &v
}

// FromBytePointer 获取*byte类型的指针的实际值，如果指针为nil的话则返回0
func FromBytePointer(p *byte) byte {
	return FromBytePointerOrDefaultIfNil(p, 0)
}

// FromBytePointerOrDefaultIfNil 获取*byte类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromBytePointerOrDefaultIfNil(v *byte, defaultValue byte) byte {
	if v == nil {
		return defaultValue
	}
	return *v
}
