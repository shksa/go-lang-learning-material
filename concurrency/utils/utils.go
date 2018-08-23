package utils

import (
	"fmt"
	"time"
)

// Timeit prints a function execution time
func Timeit(function func(int, int) float64) func(int, int) float64 {
	return func(arg1, arg2 int) float64 {
		start := time.Now()
		result := function(arg1, arg2)
		end := time.Now()
		fmt.Printf("took %f secs \n", end.Sub(start).Seconds())
		return result
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
