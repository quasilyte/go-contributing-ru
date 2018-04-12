package main

import (
	"regexp" // Regular expressions.
)

// Hint: MatchString method accepts string argument.
//       regexp.Regexp type has a method that accepts []byte argument.

func main() {
	// Normal strings use "", raw strings use ``.
	// You don't need to escape special chars in raw strings,
	// this is why they are preffered for regexp patterns.
	re := regexp.MustCompile(`\d+.\d+`)

	// If statement can introduce variables into it's scope.
	if x := "132.0"; re.MatchString("132.0") {
		println("matched:", x)
	} else {
		println("not matched:", x) // x is also reachable in "else" branch
	}

	if x := []byte("abc"); re.MatchString(x) {
		println("matched:", x)
	}
}
