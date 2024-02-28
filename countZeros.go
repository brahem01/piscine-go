package piscine

func countZeros(n int) int {
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		return 1 + countZeros(n/2)
	}
	return 1 + countZeros(n-1)
}
