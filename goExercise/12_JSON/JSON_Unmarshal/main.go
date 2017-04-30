package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string
	Last string
	Age int `json:"score"`
}

func main()  {
	var p1 person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First": "Yi", "Last": "Wen", "score": 20}`)
	json.Unmarshal(bs, &p1)

	fmt.Println("--------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}
