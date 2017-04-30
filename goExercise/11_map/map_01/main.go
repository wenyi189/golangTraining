package main

import (
	//"golang.org/x/tour/wc"
	"strings"
	"fmt"
)

func WordCount(s string) map[string]int {
	// return map[string]int{"x": 1}
	countMap := make(map[string]int)
	strs := strings.Fields(s)
	for i := 0; i < len(strs); i++ {
		countMap[strs[i]]++
	}
	return countMap
}

func main() {
	result := WordCount("I am leanring go and go is awesome")
	fmt.Println(result)
}
