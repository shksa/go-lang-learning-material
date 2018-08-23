package homepage

import (
	"fmt"
	"math/rand"
	"time"
)

func timeit(function func([][]int, [][]int, int, int) int) func([][]int, [][]int, int, int) int {
	return func(arg1, arg2 [][]int, m, n int) int {
		start := time.Now()
		result := function(arg1, arg2, m, n)
		end := time.Now()
		fmt.Printf("took %f secs \n\n\n", end.Sub(start).Seconds())
		return result
	}
}

func createMat(m, n int) [][]int {
	var mat = make([][]int, m)
	for rowIdx := range mat {
		mat[rowIdx] = make([]int, n)
		for colIdx := range mat[rowIdx] {
			mat[rowIdx][colIdx] = rand.Int()
		}
	}
	// fmt.Println(mat)
	return mat
}

func multiply(mat1 [][]int, mat2 [][]int) int {
	sum := 0
	rowsOfMat1 := len(mat1)
	colsOfMat2 := len(mat2[0])
	for mat1RowIdx := 0; mat1RowIdx < rowsOfMat1; mat1RowIdx++ {
		for mat2ColIdx := 0; mat2ColIdx < colsOfMat2; mat2ColIdx++ {
			sum += dotProduct(mat1, mat2, mat1RowIdx, mat2ColIdx)
		}
	}
	return sum
}

func dotProduct(mat1 [][]int, mat2 [][]int, mat1RowIdx int, mat2ColIdx int) int {
	result := 0

	for k := 0; k < 100; k++ {
		for idx, mat1ValueInRow := range mat1[mat1RowIdx] {
			result += mat1ValueInRow * mat2[idx][mat2ColIdx]
		}
	}
	// fmt.Println(result)
	return result
}

// CreateMatAndMultiply will create 2 matrices and return their product.
func CreateMatAndMultiply(m, n int) int {
	mat1 := createMat(m, n)
	mat2 := createMat(n, m)
	sum := multiply(mat1, mat2, m, m)
	return sum
}
