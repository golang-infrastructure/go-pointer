package pointer

// ToComplex128Pointer 将complex128类型的变量转换为对应的*complex128指针类型
func ToComplex128Pointer(v complex128) *complex128 {
	return &v
}

// ToComplex128PointerOrNilIfZero 将complex128类型的变量转换为对应的*complex128指针类型，如果变量的值为0的话则返回nil指针
func ToComplex128PointerOrNilIfZero(v complex128) *complex128 {
	if v == 0 {
		return nil
	}
	return &v
}

// FromComplex128Pointer 获取*complex128类型的指针的实际值，如果指针为nil的话则返回0
func FromComplex128Pointer(p *complex128) complex128 {
	return FromComplex128PointerOrDefaultIfNil(p, 0)
}

// FromComplex128PointerOrDefaultIfNil 获取*complex128类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromComplex128PointerOrDefaultIfNil(v *complex128, defaultValue complex128) complex128 {
	if v == nil {
		return defaultValue
	}
	return *v
}
