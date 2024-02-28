package piscine

// this function check if slice passed 
// as a parameter is sorted or not 

func isSorted(slice []int) bool {
	if len(slice) < 2 {
		return true
	}
	if slice[0] > slice[1] {
		return false
	}
	return isSorted(slice[1:])
}
