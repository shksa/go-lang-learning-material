package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

type polar struct {
	radius float64
	θ      float64
}

type cartesian struct {
	x float64
	y float64
}

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { // Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main() {
	questions := make(chan polar) // make a channel that can pass "polar" values.
	defer close(questions)
	answers := createSolver(questions) // create a channel where answers to the questions are passed.
	// At this point,
	// 2 channels are set up.
	// A seperate goroutine is setup that waits for polar values to be sent on the "questions" channel.
	// No other goroutine is blocked.
	defer close(answers)
	interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian) // make a new channel that can pass "cartesian" values

	// creating a goroutine that waits for "polar" values to be sent on the "questions" channel.
	go func() { // go statement with an anonymous function call
		fmt.Println("A seperate goroutine started")
		for { // infinite loop that waits
			polarCoord := <-questions // waits to retrive a "polar" value from the "question" channel
			θ := polarCoord.θ * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(θ)
			y := polarCoord.radius * math.Sin(θ)
			answers <- cartesian{x, y} // puts the "cartesian" value on the answers channel
		}
	}()

	return answers
}

const result = "Polar radius=%.02f θ=%.02f° → Cartesian x=%.02f y=%.02f\n"

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin) // create a reader that reads from stdin console.
	fmt.Println(prompt)

	for { // infinite loop that waits for user input on the stdin console.
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var radius, θ float64 // variables to store the user input values.
		// scans for values defined by the format string and populates the above defined variables with them.
		if _, err := fmt.Sscanf(line, "%f %f", &radius, &θ); err != nil {
			fmt.Fprintln(os.Stderr, "invalid input")
			continue
		}
		questions <- polar{radius, θ}                   // puts the "polar" value on the "questions" channel.
		coord := <-answers                              // waits to retrive a "cartesian" value from the "answers" channel.
		fmt.Printf(result, radius, θ, coord.x, coord.y) // prints the result to the console.
	}

	fmt.Println()
}
