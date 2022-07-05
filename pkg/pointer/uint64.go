package pointer

// ToUInt64Pointer 将uint64类型的变量转换为对应的*uint64指针类型
func ToUInt64Pointer(v uint64) *uint64 {
	return &v
}

// ToUInt64PointerOrNilIfZero 将uint64类型的变量转换为对应的*uint64指针类型，如果变量的值为0的话则返回nil指针
func ToUInt64PointerOrNilIfZero(v uint64) *uint64 {
	if v == 0 {
		return nil
	}
	return &v
}

// FromUInt64Pointer 获取*uint64类型的指针的实际值，如果指针为nil的话则返回0
func FromUInt64Pointer(p *uint64) uint64 {
	return FromUInt64PointerOrDefaultIfNil(p, 0)
}

// FromUInt64PointerOrDefaultIfNil 获取*uint64类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromUInt64PointerOrDefaultIfNil(v *uint64, defaultValue uint64) uint64 {
	if v == nil {
		return defaultValue
	}
	return *v
}
