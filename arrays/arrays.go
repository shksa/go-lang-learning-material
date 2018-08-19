package main

import (
	"fmt"
)

func arrayCreation() {
	// with 'var' keyword and no explicit initialization
	var buffer [10]byte
	fmt.Printf("%v\n", buffer)
	var twoDbuffer [3][3]byte
	fmt.Printf("%v\n", twoDbuffer)

	// with 'var' keyword and explicit initialization using the
	// composite type literal syntax.
	var list = [5]int{10, 11, 12, 13, 14}
	fmt.Printf("%v\n", list)
	var twoDlist = [3][3]string{
		{"a", "b", "c"},
		{"d", "e", "f"},
		{"g", "h", "i"},
	}
	fmt.Printf("%v\n", twoDlist)

	// with 'var' keyword and explicit initialization without specifying the length in the type literal.
	var cities = [...]string{"bangalore", "chennai", "mumbai"}
	fmt.Printf("%v\n", cities)
}

func main() {
	arrayCreation()
}
