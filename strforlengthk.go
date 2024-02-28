package piscine

import "fmt"

// the function print all possible string combinattion with
// a length k and with the character from the array that we
// pass as a parameter

func strOfLengthK(slice []string, n int, s string) {
	if len(s) == n {
		fmt.Println(s)
		return
	}
	for _, v := range slice {
		strOfLengthK(slice, n, s+v)
	}
}
