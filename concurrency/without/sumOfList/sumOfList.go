package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/shksa/learningGo/concurrency/utils"
)

func sum(numList []int) float64 {
	sum := float64(0)
	for _, num := range numList {
		sum += float64(num)
	}
	return sum
}

func calculateSumOfNumbersUpto(upperbound, _ int) {
	// numList := utils.NumListUpto(upperbound)
	numList := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fullSum := sum(numList)
	fmt.Printf("full sum: %f \n", fullSum)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var upperbound float64
	for {
		fmt.Printf("enter the upperbound of the list: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if _, err := fmt.Sscanf(line, "%f", &upperbound); err != nil {
			fmt.Fprintf(os.Stderr, "invalid input \n")
		}
		utils.Timeit(calculateSumOfNumbersUpto)(int(upperbound), 42)
	}
}
