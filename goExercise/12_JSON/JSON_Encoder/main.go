package main

import (
	"os"
	"encoding/json"
)

type person struct {
	First string
	Last string
	Age int
	notExported int
}

func main() {
	p1 := person{"Yi", "Wen", 26, 100}
	json.NewEncoder(os.Stdout).Encode(p1)
}