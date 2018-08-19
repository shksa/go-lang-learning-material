package main

import (
	"fmt"
)

func typeConversions() {
	stringVar := "sleptking"
	// converting string type to byte slice type []byte
	// using the call Type(expression)
	byteSlice := []byte(stringVar)
	fmt.Printf("%v \n", byteSlice) // [115 108 101 112 116 107 105 110 103]

	// converting string type to rune slice type []byte
	// using the call Type(expression)
	runeSlice := []rune(stringVar)
	fmt.Printf("%v \n", runeSlice) // [115 108 101 112 116 107 105 110 103]
}

func main() {
	typeConversions()
}
