package main

import (
	"fmt"

	"github.com/shksa/learningGo/concurrency/utils"
)

func sum(partialNumList []int, partialSums chan float64) {
	sum := float64(0)
	for _, num := range partialNumList {
		sum += float64(num)
	}
	fmt.Println(sum)
	partialSums <- sum // send sum to the "partialSums" channel
}

func addPartialSums(partialSumsCh chan float64, noOfGoroutines int) float64 {
	var fullSum float64
	for idx := 0; idx < noOfGoroutines; idx++ {
		fullSum += <-partialSumsCh
	}
	return fullSum
}

func calculateSumOfNumbersUpto(upperbound, noOfGoroutines int) {
	numList := utils.NumListUpto(upperbound)
	partialSumsCh := make(chan float64, noOfGoroutines)

	listSize := len(numList)
	partialListSize := listSize / noOfGoroutines

	var startIdx int

	for startIdx = 0; (listSize - startIdx) >= partialListSize; startIdx += partialListSize {
		go sum(numList[startIdx:startIdx+partialListSize], partialSumsCh)
	}
	go sum(numList[startIdx:], partialSumsCh)

	fullSum := addPartialSums(partialSumsCh, noOfGoroutines)

	fmt.Printf("full sum: %f \n", fullSum)
}

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// var upperbound float64
	// var noOfGoroutines int
	// for {
	// 	fmt.Printf("enter the upperbound of the list and the no. of goroutines. ex: 1e5 6, 1e4 3 :")
	// 	line, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		break
	// 	}
	// 	if _, err := fmt.Sscanf(line, "%f %d", &upperbound, &noOfGoroutines); err != nil {
	// 		fmt.Fprintf(os.Stderr, "invalid input \n")
	// 	}
	// 	utils.Timeit(calculateSumOfNumbersUpto)(int(upperbound), noOfGoroutines)
	// }

}
