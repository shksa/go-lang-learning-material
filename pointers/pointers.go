package main

import "fmt"

func decorateProgram(title string, fn func()) func() {
	return func() {
		println("\n\n--------------------------------------------------")
		println(title)
		println("--------------------------------------------------\n\n")
		fn()
		println("--------------------------------------------------\n\n")
	}
}

var decorateExample = (func() func(string, func()) func() {
	var exampleNumber = 0
	return func(exampleTitle string, fn func()) func() {
		return func() {
			exampleNumber++
			println("	--------------------------------------------------")
			fmt.Printf("ex: %d	%s\n", exampleNumber, exampleTitle)
			println("	--------------------------------------------------")
			fn()
			println("	--------------------------------------------------\n\n")
		}
	}
})()

func swap(px, py *int) {
	println("	swap fn start\n")
	println("		px =", px, ", py =", py)
	println("		before swap, *px =", *px, ", *py =", *py)
	*px, *py = *py, *px
	println("		after swap, *px =", *px, ", *py =", *py)
	println("\n	swap fn end")
}

// Vertex is a type which inherits the built-in structure type
type Vertex struct {
	X, Y int
}

// This function receives a copy of the argument's value as it's parameter's value
func scaleV1(v Vertex, scaleParam int) {
	v.X = v.X * scaleParam
	v.Y = v.Y * scaleParam
}

// This function receives the address of the argument as it's parameter's value
// i.e The parameter becomes a pointer to the argument variable.
// Though this pointer the argument variable's value can be changed, this is called
// indirecting.
func scaleV2(pv *Vertex, scaleParam int) {
	println("	scaleV2 func start\n")
	println("		pv =", pv)
	println("		before scaling, (*pv).X =", (*pv).X, ", (*pv).Y =", (*pv).Y)
	(*pv).X = (*pv).X * scaleParam
	(*pv).Y = (*pv).Y * scaleParam
	println("		after scaling, *pv).X =", (*pv).X, ", (*pv).Y =", (*pv).Y)
	println("\n	scaleV2 func end")
}

var swapTwoVariables = decorateExample(
	"Swapping values of two variables through their pointers",
	func() {
		x, y := 1, 2
		println("	before swap fn call, x =", x, ", y =", y)
		swap(&x, &y)
		println("	after swap fn call, x =", x, ", y =", y)
	},
)

var modifyingStructFields = decorateExample(
	"Changing the fields of a struct type variable through it's pointer",
	func() {
		v := Vertex{2, 4}
		println("	Vertex v before scaling, v.x =", v.X, ", v.y =", v.Y)
		scaleV2(&v, 10)
		println("	Vertex v after scaling, v.x =", v.X, ", v.y =", v.Y)
	},
)

var creatingValuesAndPointersAtOnce = decorateExample(
	"Creating values and pointers to them at the same time using new() anf &",
	func() {
		var pi = new(int)
		fmt.Printf("	pi, typeof pi, *pi, typeof *pi, %T, %v, %T, %v\n", pi, pi, *pi, *pi)
	},
)

var runExamples = decorateProgram(
	"This is how pointers work in Go",
	func() {
		swapTwoVariables()
		modifyingStructFields()
		creatingValuesAndPointersAtOnce()
	},
)

func main() {
	runExamples()
}
