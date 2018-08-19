package main

import (
	"fmt"
)

// Vertex is customized structure type that represents a vertex coordinate
type Vertex struct{ x, y int }

// Scale scales values of Vertex type
func (v Vertex) Scale(f int) {
	v.x = v.x * f
	v.y = v.y * f
}

// Move moves values of Vertex type
func (v Vertex) Move(dx, dy int) {
	v.x = v.x + dx
	v.y = v.y + dy
}

func main() {
	var v = Vertex{3, 4}
	fmt.Printf("Vertex value before scaling: %#v \n", v)

	// scale v with it's method scale()
	v.Scale(2)
	fmt.Printf("Vertex v after scaling by it's method Scale(2): %#v \n", v)

	// move v with it's method move()
	v.Move(5, 5)
	fmt.Printf("Vertex v after moving by it's method Move(3): %#v \n", v)

	// scale v by calling scale() method on it's pointer
	(&v).Scale(4)
	fmt.Printf("Vertex v after scaling by calling Scale(3) on it's pointer: %#v \n", v)

	// move v by calling move() method on it's pointer
	(&v).Move(4, 4)
	fmt.Printf("Vertex v after scaling by calling Move(3) on it's pointer: %#v \n", v)
}
