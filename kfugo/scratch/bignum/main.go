package main

import (
	"fmt"      // Formatted I/O with functions analogous to C's printf and scanf.
	"log"      // Simple logging package
	"math/big" // Arbitrary-precision arithmetic (big numbers).
)

// Hint: ":=" declares and initializes, but "=" only assigns to already defined names.

func main() {
	x := new(big.Int) // new returns *T, in this case, it's *big.Int
	y := &big.Int{}   // It's also possible to use literal with "&"

	x, ok = x.SetString("999888777666555444333222111000", 10)
	if !ok {
		log.Fatal("failed to set value for x")
	}
	// Can use "_" to ignore certain return value.
	y, _ = y.SetString("1", 10)

	var result big.Int // Defines result variable. Initialized to zero value.
	result.Add(x, y)

	fmt.Printf("result=%s\n", result.String())
}
