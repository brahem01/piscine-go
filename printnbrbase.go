package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	str := ""
	negative := false
	if nbr < 0 {
		nbr = -nbr
		negative = true
	}
	for nbr != 0 {
		str = string(base[nbr%len(base)]) + str
		nbr /= len(base)
	}
	if negative {
		str = "-" + str
	}
	for _, char := range str {
		z01.PrintRune(char)
	}
}
