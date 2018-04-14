package main

import (
	"log"     // Simple logging package.
	"strconv" // Conversions to and from string representations.
)

// Tip: explicit type conversion (cast) is needed to compare
//      two values of different types.

func main() {
	s := "1234"
	x, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("atoi failed: %v", err)
	}

	// ParseInt can be used instead of atoi to parse
	// numbers of different bases and widths.
	base := int(10)
	bitSize := int(64)
	y, err := strconv.ParseInt(s, base, bitSize)
	if err != nil {
		log.Fatalf("parse int failed: %v", err)
	}

	if y == x {
		println("x and y are equal")
	}
}
