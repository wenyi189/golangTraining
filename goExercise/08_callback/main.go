package main

import	"fmt"

func filter(number []int, callback func(int) bool) []int {
	var slice []int
	for _, item := range number {
		if callback(item) {
			slice = append(slice, item)
		}
	}
	return slice
}

func main()  {
	slice := filter([]int{10, 20, 30, 35, 40}, func(n int) bool {
		return n > 20
	})
	fmt.Println(slice)
}
