package main

import (
	"fmt"
)

func main() {
	var i interface{} = 99
	var s interface{} = []string{"bob", "dylan"}
	/* Note that if we printed the original i and s variables (both of type interface{})
	they would be printed as an int and a []string. This is because when the
	fmt package’s print functions are faced with interface{} types,
	they are sensible enough to print the actual underlying values. */

	if i, ok := i.(int); ok {
		fmt.Printf("%T→%d\n", i, i) // i is a shadow variable of type int
		// int→99
	}

	if s, ok := s.([]string); ok {
		fmt.Printf("%T→%q\n", s, s) // s is a shadow variable of type []string
		// []string→["bob" "dylan"]
	}
}
