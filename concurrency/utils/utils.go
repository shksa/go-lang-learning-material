package utils

import (
	"fmt"
	"time"
)

// Timeit prints a function execution time
func Timeit(function func(int, int)) func(int, int) {
	return func(arg1, arg2 int) {
		start := time.Now()
		function(arg1, arg2)
		end := time.Now()
		fmt.Printf("took %f secs \n\n\n", end.Sub(start).Seconds())
	}
}

// NumListUpto returns a []int
func NumListUpto(upperbound int) []int {
	numList := make([]int, upperbound)
	for idx := range numList {
		numList[idx] = idx
	}
	return numList
}
