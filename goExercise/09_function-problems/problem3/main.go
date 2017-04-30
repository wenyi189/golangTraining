//Write a function with one variadic parameter that finds the greatest number in a list of numbers.

package main

import "fmt"

func findGreat(n ...int) int {
	var result int
	for _, v := range n {
		if v > result {
			result = v
		}
	}
	return result
}

func main() {
	fmt.Println(findGreat(21, 32, 541, 1, 32, 9393, 21))
}

