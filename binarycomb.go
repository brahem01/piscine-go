package piscine

// the function print all the possible binary combination
// with length n, with a condition say:
// the number of 1 in the left side equal the right side
import (
	"fmt"
)

func recu(s string, sum, sum2, n int) {
	if len(s) == 2*n {
		if sum == sum2 {
			fmt.Println(s)
		}
		return
	}
	recu(s+"0", sum, sum2, n)
	if len(s) < n {
		recu(s+"1", sum+1, sum2, n)
	} else {
		recu(s+"1", sum, sum2+1, n)
	}
}
