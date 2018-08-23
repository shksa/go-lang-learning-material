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
	interact()
}

func solve(polarCoord polar) cartesian {
	θ := polarCoord.θ * math.Pi / 180.0
	x := polarCoord.radius * math.Cos(θ)
	y := polarCoord.radius * math.Sin(θ)
	return cartesian{x, y}
}

const result = "Polar radius=%.02f θ=%.02f° → Cartesian x=%.02f y=%.02f\n"

func interact() {
	reader := bufio.NewReader(os.Stdin) // create a reader that reads from stdin console.
	fmt.Println(prompt)

	var polarCoordFromUser polar // variables to store the user input values.
	for {                        // infinite loop that waits for user input on the stdin console.
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// scans for values defined by the format string and populates the above defined variables with them.
		if _, err := fmt.Sscanf(line, "%f %f", &polarCoordFromUser.radius, &polarCoordFromUser.θ); err != nil {
			fmt.Fprintln(os.Stderr, "invalid input")
			continue
		}
		cartesianCoord := solve(polarCoordFromUser)
		fmt.Printf(result, polarCoordFromUser.radius, polarCoordFromUser.θ, cartesianCoord.x, cartesianCoord.y) // prints the result to the console.
	}

	fmt.Println()
}
