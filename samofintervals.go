package piscine

func SumOfIntervals(intervals [][2]int) (n int) {
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
