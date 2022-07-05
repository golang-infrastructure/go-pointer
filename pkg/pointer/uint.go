package pointer

// ToUIntPointer 将uint类型的变量转换为对应的*uint指针类型
func ToUIntPointer(v uint) *uint {
	return &v
}

// ToUIntPointerOrNilIfZero 将uint类型的变量转换为对应的*uint指针类型，如果变量的值为0的话则返回nil指针
func ToUIntPointerOrNilIfZero(v uint) *uint {
	if v == 0 {
		return nil
	}
	return &v
}

// FromUIntPointer 获取*uint类型的指针的实际值，如果指针为nil的话则返回0
func FromUIntPointer(p *uint) uint {
	return FromUIntPointerOrDefaultIfNil(p, 0)
}

// FromUIntPointerOrDefaultIfNil 获取*uint类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromUIntPointerOrDefaultIfNil(v *uint, defaultValue uint) uint {
	if v == nil {
		return defaultValue
	}
	return *v
}
