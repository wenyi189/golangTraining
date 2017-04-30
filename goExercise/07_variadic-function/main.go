package main

import	"fmt"

func main()  {
	n := average(43, 56, 81, 12, 45, 59)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}