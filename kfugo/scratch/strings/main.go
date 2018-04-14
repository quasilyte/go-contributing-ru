// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt" // Formatted I/O with functions analogous to C's printf and scanf.
)

// Hint: Go strings are immutable.

func main() {
	s := "I❤Go" // Go string literals are in utf8.

	// For/range loop over string yields 4-byte (int32) wide "runes".
	// The [x:y] notation creates a substring (in Go it's called slicing).
	for i, c := range s[1:] {
		fmt.Printf("index=%d type=%T, value='%c'\n", i, c, c)
	}

	s[0] = 'a'

	// Use simple for to iterate over string individual bytes.
	for i := 0; i < len(s); i++ {
		c := s[i]
		fmt.Printf("index=%d type=%T, value='%c'\n", i, c, c)
	}
}
