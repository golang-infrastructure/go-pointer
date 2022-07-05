package pointer

// ToRunePointer 将rune类型的变量转换为对应的*rune指针类型
func ToRunePointer(v rune) *rune {
	return &v
}

// ToRunePointerOrNilIfZero 将rune类型的变量转换为对应的*rune指针类型，如果变量的值为0的话则返回nil指针
func ToRunePointerOrNilIfZero(v rune) *rune {
	if v == 0 {
		return nil
	}
	return &v
}

// FromRunePointer 获取*rune类型的指针的实际值，如果指针为nil的话则返回0
func FromRunePointer(p *rune) rune {
	return FromRunePointerOrDefaultIfNil(p, 0)
}

// FromRunePointerOrDefaultIfNil 获取*rune类型的指针的实际值，如果指针为nil的话则返回defaultValue
func FromRunePointerOrDefaultIfNil(v *rune, defaultValue rune) rune {
	if v == nil {
		return defaultValue
	}
	return *v
}

