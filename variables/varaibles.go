package main

import (
	"fmt"
)

func variableCreation() {
	// single variable declaration
	var x1 int = 9 // 1st alternative

	var x2 int // 2nd alternative
	x2 = 9

	var x3 = 9 // 3rd alternative

	x4 := 9 // 4th alternative
	fmt.Println(x1, x2, x3, x4)

	// Multi variable declaration
	a, b, c := 1, 2, 3 // with short variable declaration syntax
	fmt.Println(a, b, c)

	var x, y, z int // with 'var' declaration and no explicit initialization
	fmt.Println(x, y, z)

	var m, l, k = 1, 2, 3 // with 'var' declaraation and explicit initialization
	fmt.Println(m, l, k)
}

func main() {
	variableCreation()
}
