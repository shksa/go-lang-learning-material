package main

import (
	"fmt"
)

func sliceCreation() {
	// using var declaration without explicit initialization
	var sliceA []int
	fmt.Printf("%v %v %v \n", sliceA, len(sliceA), cap(sliceA)) // [] 0 0
	fmt.Printf("%v \n", sliceA == nil)                          // true
	fmt.Printf("%T \n", sliceA)                                 // []int

	// using var declaration and explicit initialization
	var sliceB = []int{2, 4, 6}
	fmt.Printf("%v \n", sliceB) // [2, 4, 6]
	fmt.Printf("%T \n", sliceB) // []int

	// using var edclaration and explicit initialization with empty array
	var sliceC = []int{}
	fmt.Printf("%v %v %v \n", sliceC, len(sliceC), cap(sliceC)) // [] 0 0
	fmt.Printf("%v \n", sliceC != nil)                          // true

	// using make(Type, length, capacity)
	var sliceD = make([]bool, 4, 10)
	fmt.Printf("%v \n", sliceD) // [false false false false]
	fmt.Printf("%T \n", sliceD) // []bool
}

func main() {
	sliceCreation()
}
