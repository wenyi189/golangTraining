package main

import "fmt"

func main() {
	var myName string
	fmt.Print("Please enter your name:")
	fmt.Scan(&myName)
	fmt.Println("Hello, ", myName, "The name is ", myName)
}

