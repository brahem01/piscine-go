package piscine

import "math/big"

// in this function we get the square and the cube for in number and reverse him
// if the reverses are prime number we count 1
// the first numbe is 89:
// => the square of 89 is :7921 and the reverse is 1279 (prime number)
// => the cube of 89 is : 704969 and the reverse is 969407 (prime number)
// so if we pass 1 as a parameter the output is 89

func SqCubRevPrime(n int) int {
	counter := 0
	res := 89
	for counter != n {
		if big.NewInt(int64(revNum(res*res))).ProbablyPrime(20) && big.NewInt(int64(revNum(res*res*res))).ProbablyPrime(20) {
			counter++
		}
		res++
	}
	return res - 1
}

func revNum(n int) int {
	nbr := 0
	for n != 0 {
		digit := n % 10
		nbr = nbr*10 + digit
		n /= 10
	}
	return nbr
}
