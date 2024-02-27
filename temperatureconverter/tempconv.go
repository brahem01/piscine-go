package main 

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	from:= os.Args[1]
	to := os.Args[3]
	n:= atoi(os.Args[2])
	switch from {
	case "f":
		switch to{
		case "c":
			FtoC(n)
		case "k":
			FtoK(n)
		default:
			fmt.Println("invalid temperature type..!")
			return
		}
	case "k":
		switch to{
		case "c":
			KtoC(n)
		case "f":
			KtoF(n)
		default:
			fmt.Println("invalid temperature type..!")
			return
		}	
	case "c":
		switch to{
		case "f":
			CtoF(n)
		case "k":
			CtoK(n)
		default:
			fmt.Println("invalid temperature type..!")
			return
		}	
	default:
		fmt.Println("inValid temperature type !")
		return
	}
}
func FtoC(n int)  {
	fmt.Println((float64(n)-32)*5/9)
}

func FtoK(n int)  {
	fmt.Println((float64(n) - 32) * 5 / 9 + 273.15)
}

func KtoC(n int)  {
	fmt.Println(float64(n)-273.15)
}

func KtoF(n int)  {
	fmt.Println((float64(n) - 273.15) * 9 / 5 + 32)
}

func CtoF(n int)  {
	fmt.Println(float64(n)*9/5+32)
}

func CtoK(n int)  {
	fmt.Println((float64(n)+273.15))
}

func atoi(s string) (n int) {
	negative := false
	for i,v := range s {
		if i == 0 && v == '-'{
			negative = true
		}else if i == 0 && v == '+'{
			negative = false
		}else if v >= '0' && v <= '9'{
			digit:= int(v - 48)
			n = n*10 + digit
		}else{
			return 
		}
	}
	if negative{
		return -n
	}
	return n
}