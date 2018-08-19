package main

import (
	"strings"
)

// Pic returns a 2D slice of uint8 numbers that represents pixel values
func Pic(dy, dx int) [][]uint8 {
	// var sliceToReturn = make([][]uint8, dy)
	// for y := range sliceToReturn {
	// 	sliceToReturn[y] = make([]uint8, dx)
	// 	for x := range sliceToReturn[y] {
	// 		sliceToReturn[y][x] = uint8(x * y)
	// 	}
	// }

	// var sliceToReturn = make([][]uint8, dy)
	// for y := range sliceToReturn {
	// 	for x := 0; x < dx; x++ {
	// 		sliceToReturn[y] = append(sliceToReturn[y], uint8(x*y))
	// 	}
	// }

	// var sliceToReturn [][]uint8
	// for y := 0; y < dy; y++ {
	// 	sliceToReturn = append(sliceToReturn, make([]uint8, 0))
	// 	for x := 0; x < dx; x++ {
	// 		sliceToReturn[y] = append(sliceToReturn[y], uint8(x*y))
	// 	}
	// }

	var sliceToReturn [][]uint8
	for y := 0; y < dy; y++ {
		sliceToReturn = append(sliceToReturn, make([]uint8, dx))
		for x := range sliceToReturn[y] {
			sliceToReturn[y][x] = uint8(x * y)
		}
	}
	println(sliceToReturn)
	return sliceToReturn
}

// WordCount returns a map of word counts of a string
func WordCount(s string) map[string]int {
	// mapToReturn := make(map[string]int)
	var mapToReturn map[string]int = make(map[string]int)
	println("mapToReturn = ", mapToReturn)
	words := strings.Fields(s)
	for _, word := range words {
		mapToReturn[string(word)]++
	}
	return mapToReturn
}

func uninitSliceWorking() {
	println("	--------------------------------------------------")
	println("ex:2	Slices that are only declared without initializing")
	println("	--------------------------------------------------")
	var sliceOfIntegerSlices [][]int
	println("	var sliceOfIntegerSlices [][]int\n")
	println("	sliceOfIntegerSlices value = ", sliceOfIntegerSlices)
	sliceOfIntegerSlices = append(sliceOfIntegerSlices, []int{3, 4, 5}, []int{13, 14, 15})
	println("	sliceOfIntegerSlices value after appending with append() = ", sliceOfIntegerSlices)
	for i := range sliceOfIntegerSlices {
		println("\n	sliceOfIntegerSlices[", i, "] = ", sliceOfIntegerSlices[i], "\n")
		for j, intValue := range sliceOfIntegerSlices[i] {
			println("		sliceOfIntegerSlices[", i, "][", j, "] = ", intValue)
		}
	}
}

func initSliceWorking() {
	println("	--------------------------------------------------")
	println("ex:1	Slices initialized along with their declaration")
	println("	--------------------------------------------------")
	var testSlice = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}
	println(`	var testSlice = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}`, "\n")
	println("	testSlice value = ", testSlice)
	for i := range testSlice {
		println("\n	testSlice[", i, "] value = ", testSlice[i], "\n")
		for j, intValue := range testSlice[i] {
			println("		testSlice[", i, "][", j, "] = ", intValue)
		}
	}
	println("	--------------------------------------------------\n\n")
}

func sliceWorking() {
	println("\n\n--------------------------------------------------")
	println("This is how slices work in Go")
	println("--------------------------------------------------\n\n")
	initSliceWorking()
	uninitSliceWorking()
	println("\n\n--------------------------------------------------\n\n")
}

func variablesWorking() {
	println("\n\n--------------------------------------------------")
	println("This is how variables work in Go")
	println("--------------------------------------------------\n\n")
	var x int
	println("value of x declared as 'var x int' is ", x)
	x = 99
	println("value of x after changing it's value by assignment statement 'x = 99' is ", x)
	var y = `foo` // type of y is infered by the type of value on rhs i.e type of `foo`
	println("value of y declared as 'var y = `foo`' is ", y)
}

func enumWorking() {
	const (
		goerge = iota // value is 0
		paul          // also iota but now iota's value is 1
		ringo         // also iota but now iota's value is 2
		lenon         // also iota but now iota's value is 3
	)
	println("with iota initialization: goerge, paul, ringo, lenon = ", goerge, paul, ringo, lenon)
	const (
		metallica    = 0
		blackSabbath // 0
		megaDeth     // 0
		ozzy         // 0
	)
	println("without iota initialization: metallica, blackSabbath, megaDeth, ozzy", metallica, blackSabbath, megaDeth, ozzy)
}

func main() {
	// pointerWorking()
	// println("\n******************************************************************************\n")
	// sliceWorking()
	// println("\n******************************************************************************\n")
	// variablesWorking()
	// Pic(5, 5)
	// pic.Show(Pic)
	// wc.Test(WordCount)
	// enumWorking()
}
