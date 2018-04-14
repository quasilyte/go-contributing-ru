// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"  // Formatted I/O with functions analogous to C's printf and scanf.
	"sort" // Functions for sorting
)

// Hint: ":=" is only valid inside function scope.
//       It can't be used at top level (outside functions).

lucky := 777

func main() {
	xs := []int{4, 5, 1, 3, lucky}
	sort.Ints(xs)
	fmt.Printf("xs=%v\n", xs)

	var ys = []float32{4.4, 14.55, 0, 1}
	// Sort by using sorting predicate for "less".
	// This sorts in ASC order.
	sort.Slice(ys, func(i, j int) bool {
		return ys[i] < ys[j]
	})
	fmt.Printf("ys=%v\n", ys)
	// This sorts in DESC order.
	sort.Slice(ys, func(i, j int) bool {
		return ys[i] > ys[j]
	})
	fmt.Printf("ys=%v\n", ys)
}
