package main

import (
	"fmt"
	"strings"
)

func main() {
	xs := []string{"Hello", "World"}
	fmt.Println(strings.Join(xs, ", "))
}
