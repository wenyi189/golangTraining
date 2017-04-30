package main

import (
	"encoding/json"
	"strings"
	"fmt"
)

type person struct {
	First string
	Last string
	Age int
	notExported int
}

func main() {
	var p1 person
	reader := strings.NewReader(`{"First":"Yi","Last":"Wen","Age":26}`)
	json.NewDecoder(reader).Decode(&p1)

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}