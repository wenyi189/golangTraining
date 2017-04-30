package main

import (
	"fmt"
	"sort"
)

func main() {
	studyGroup := []string{"Zeno", "John", "Al", "Jenny", "Yi"}

	fmt.Println(studyGroup)
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println(studyGroup)
}
