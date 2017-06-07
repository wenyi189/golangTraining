package main

import (
	"fmt"
)

//func main() {
//	c := make(chan int)
//	c <- 1
//	fmt.Println(<-c)
//}

// This results in a deadlock
// can you determine why?
// and what would you do to fix it?

func main() {
	c := make(chan int)
	go func () {
		c <- 1
	}()
	fmt.Println(<-c)
}