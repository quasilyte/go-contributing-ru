package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Point описывает точку в двухмерном пространстве.
type Point struct {
	X float64
	Y float64
}

func main() {
	a := Point{X: 1, Y: 2}

	data, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v => %s\n", a, data)
	// Печатает:
	// {1 2} => {"X":1,"Y":2}
}
