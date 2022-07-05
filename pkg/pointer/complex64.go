package pointer

// ToComplex64Pointer 将complex64类型的变量转换为对应的*complex64指针类型
func ToComplex64Pointer(v complex64) *complex64 {
	return &v
}

// ToComplex64PointerOrNilIfZero 将complex64类型的变量转换为对应的*complex64指针类型，如果变量的值为0的话则返回nil指针
func ToComplex64PointerOrNilIfZero(v complex64) *complex64 {
	if v == 0 {
		return nil
	}
	return &v
}

// FromComplex64Pointer 获取*complex64类型的指针的实际值，如果指针为nil的话则返回0
func FromComplex64Pointer(p *complex64) complex64 {
	return FromComplex64PointerOrDefaultIfNil(p, 0)
}

// FromComplex64PointerOrDefaultIfNil 获取*complex64类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromComplex64PointerOrDefaultIfNil(v *complex64, defaultValue complex64) complex64 {
	if v == nil {
		return defaultValue
	}
	return *v
}
