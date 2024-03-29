package piscine

// this function check if slice passed
// as a parameter is sorted or not
func isSorted(slice []int) bool {
	return helper(slice, len(slice)-1)
}

func helper(slice []int, count int) bool {
	if count == 1 {
		return true
	}
	return slice[count] > slice[count-1] && helper(slice, count-1)
}
