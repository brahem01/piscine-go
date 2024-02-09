package piscine

func ItoaBase(value, base int) string {
	Base := "0123456789abcdef"
	str := ""
	negative := false
	if value < 0 {
		value = -value
		negative = true
	}
	for value != 0 {
		str = string(Base[value%base]) + str
		value/=base
	}
	if negative {
		return "-" + str
	}
	return str

}