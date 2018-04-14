package main

import (
	"fmt"       // Formatted I/O with functions analogous to C's printf and scanf.
	"io/ioutil" // Convenient functions for simple I/O tasks.
	"log"       // Simple logging package.
	"os"        // OS-dependent functions.
)

// Hint: string can be converted to []byte using type conversion.

func main() {
	// Type conversion from untyped literal to os.FileMode:
	permissions := os.FileMode(0666)
	filename := "my_file.txt"
	err := ioutil.WriteFile(filename, "123", permissions)
	if err != nil {
		log.Fatalf("write file: %v", err)
	}

	// Read file contents.
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("read file: %v", err)
	}
	fmt.Printf("data: '%s'\n", data)

	if err := os.Remove(filename); err != nil {
		log.Fatalf("remove file: %v", err)
	}
}
