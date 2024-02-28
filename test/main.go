package main

import (
	"fmt"
)
var arr = []int{}
func main() {
	slice := []int{7,20,3,7,5,6,7,8}
	fmt.Println(linearSearch(slice, 7))
}

func linearSearch(slice []int, target int) []int {
	return helper(slice, target, len(slice)-1)
}

func helper(slice []int, target, count int) []int {
	if count < 0 {
		return nil
	}
	if slice[count] == target {
		return  append(helper(slice, target, count-1), count)
	}
	return helper(slice, target, count-1)
}

// func helper(slice []int, target, count int) []int {
// 	if count == len(slice) {
// 		return arr
// 	}
// 	if slice[count] ==  target {
// 		arr = append(arr, count)
// 	}
// 	return  helper(slice, target, count+1)
// }
