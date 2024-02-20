package main

import (
	"fmt"
)

func main() {
	fmt.Println(SumOfIntervals([][2]int{{1, 6}, {-7, 10}, {3, 5}}))
}


func SumOfIntervals(intervals [][2]int) (n int) {
	for i := 0; i< len(intervals); i++ {
		for j:= 0; j < len(intervals)- i -1; j++{
			if intervals[j][0] > intervals[j+1][0]{
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			}
		}
	}
	fmt.Println(intervals)
	num := intervals[0][0]
	check := make(map[[2]int]bool)
	for _, v := range intervals {
	  if num > v[0]{
		v[0] = num
	  }
	  if !check[v]  {
		n += (v[1]-v[0])
		check[v] = true
		num = v[1]
	  }
	}
	return 
  }

