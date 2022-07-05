package pointer

// ToFloat64Pointer 将float64类型的变量转换为对应的*float64指针类型
func ToFloat64Pointer(v float64) *float64 {
	return &v
}

// ToFloat64PointerOrNilIfZero 将float64类型的变量转换为对应的*float64指针类型，如果变量的值为0的话则返回nil指针
func ToFloat64PointerOrNilIfZero(v float64) *float64 {
	if v < 0.0000001 {
		return nil
	}
	return &v
}

// FromFloat64Pointer 获取*float64类型的指针的实际值，如果指针为nil的话则返回0
func FromFloat64Pointer(p *float64) float64 {
	return FromFloat64PointerOrDefaultIfNil(p, 0)
}

// FromFloat64PointerOrDefaultIfNil 获取*float64类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromFloat64PointerOrDefaultIfNil(v *float64, defaultValue float64) float64 {
	if v == nil {
		return defaultValue
	}
	return *v
}
