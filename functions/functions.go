package main

import (
	"fmt"
)

func createFunctions() {
	// with "var" keyword and explicit initialization and specifying the function type
	var sum func(int, int) int = func(a, b int) int { return a + b }
	fmt.Printf("Type: %T, Value: %v, Call(2, 3): %v \n", sum, sum, sum(2, 3))
}

func main() {
	createFunctions()
}
