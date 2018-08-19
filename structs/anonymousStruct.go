package main

import "fmt"

func createAnonymousStructs() {

	// with 'var' declaration and NO explicit initialization
	var name struct{ first, last string }
	name.first = "andres"
	name.last = "iniesta"
	fmt.Printf("%#v \n", name)

	// SVD with struct type spec and composite literal
	// ex: 1
	person := struct {
		age        int
		name       string
		isMarried  bool
		ownsABitch bool
	}{
		name: "lionel messi",
		age:  20,
	}
	fmt.Printf("%v \n", person)
	// {20 lionel messi false false}

	// ex: 2 // by not including the field names in composite literal
	points := []struct{ x, y int }{
		{4, 6},
		{},
		{-7, 11},
		{15, 17},
		{4, 8},
	}
	for _, point := range points {
		fmt.Printf("(%d, %d) \n", point.x, point.y)
	}
	/*
		(4, 6)
		(0, 0)
		(-7, 11)
		(15, 17)
		(4, 8)
	*/

}

func main() {
	createAnonymousStructs()
	EmbedStructs()
}
