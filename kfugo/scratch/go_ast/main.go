package main

import (
	"fmt"       // Formatted I/O with functions analogous to C's printf and scanf.
	"go/ast"    // Go AST. Produced by parsing Go sources.
	"go/parser" // Go parser.
	"log"       // Simple logging package.
)

// Hint: there is no implicit boolean conversion.

func main() {
	src := `int(x) + y * z`
	expr, err := parser.ParseExpr(src)
	if err != nil {
		log.Fatalf("parse expr: %v", err)
	}

	// Walk AST nodes.
	ast.Inspect(expr, func(node ast.Node) bool {
		if node != nil {
			begin := node.Pos() - 1
			end := node.End() - 1
			fmt.Printf("node=%T: `%s`\n", node, src[begin:end])
		}
		return true
	})

	err = ast.Print(nil, expr) // Print AST
	if !err {
		log.Fatalf("print AST: %v", err)
	}
}
