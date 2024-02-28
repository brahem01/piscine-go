package piscine

// the function return the sum of the digit of number
// passage as a parameter, if the sum > 10, the function
// call her self

func DigitalRoot(n int) int {
	// ...
	if n < 10 {
		return n
	}
	s := itoa(n)
	n = 0
	for _, v := range s {
		n += int(v - 48)
	}
	return DigitalRoot(n)
}
func itoa(n int) string {
	s := ""
	for n != 0 {
		char := rune(n%10 + 48)
		s += string(char)
		n /= 10
	}
	return s
}
