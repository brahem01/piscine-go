package piscine

// the expsum is a function that
// return the number of combination
// of sums of numbers that can equal a
// number n type uint64

func expSum(n uint64) uint64 {
	partitions := make([]uint64, n+1)
	partitions[0] = 1
	for i := uint64(1); i <= n; i++ {
		for j := i; j <= n; j++ {
			partitions[j] += partitions[j-i]
		}
	}

	return partitions[n]
}
