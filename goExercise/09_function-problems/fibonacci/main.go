package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	x := 0
	y := 1
	return	func() int {
		x,y = y, x+y
		fmt.Println("x:", x, "y:", y)
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
