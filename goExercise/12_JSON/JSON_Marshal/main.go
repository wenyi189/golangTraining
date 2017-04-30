package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string
	Last string
	Age int
	notExported int
}

func main()  {
	p1 := person{"Yi", "Wen", 26, 100}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}
