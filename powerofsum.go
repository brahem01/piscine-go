package piscine

func powerOfSum(x, n, i, sum int, count *int)  {
	if i == 50 {
		if sum == x {
			*count++
		}
		return 
	}
	if sum + power(i, n) <= x{
		powerOfSum(x, n, i+1, sum+ power(i, n), count)
	}
	powerOfSum(x, n, i+1, sum, count) 
}

func power(n, Power int) int {
	result := 1
	for Power != 0 {
		result *=n
		Power--
	}
	return result 
}