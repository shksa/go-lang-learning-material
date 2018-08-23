package main

import (
	"fmt"
	"math/rand"

	"github.com/shksa/learningGo/concurrency/utils"
)

func createMat(noOfRows, noOfCols int) [][]float64 {
	var mat = make([][]float64, noOfRows)
	for rowIdx := range mat {
		mat[rowIdx] = make([]float64, noOfCols)
		for colIdx := range mat[rowIdx] {
			mat[rowIdx][colIdx] = rand.Float64() * 1e3
		}
	}
	// fmt.Println(mat)
	return mat
}

func matrixMultiply(mat1 [][]float64, mat2 [][]float64) float64 {
	var sum float64

	rowsOfMat1 := len(mat1)
	colsOfMat2 := len(mat2[0])
	for mat1RowIdx := 0; mat1RowIdx < rowsOfMat1; mat1RowIdx++ {
		for mat2ColIdx := 0; mat2ColIdx < colsOfMat2; mat2ColIdx++ {
			sum += dotProduct(mat1, mat2, mat1RowIdx, mat2ColIdx)
		}
	}

	return sum
}

func dotProduct(mat1 [][]float64, mat2 [][]float64, mat1RowIdx int, mat2ColIdx int) float64 {
	var result float64

	for k := 0; k < 100; k++ {
		for mat1ColIdx := range mat1[mat1RowIdx] {
			result += mat1[mat1RowIdx][mat1ColIdx] * mat2[mat1ColIdx][mat2ColIdx]
		}
	}
	// fmt.Println(result)
	return result
}

func createMatAndMultiply(m, n int) float64 {
	mat1 := createMat(m, n)
	mat2 := createMat(n, m)
	sum := matrixMultiply(mat1, mat2)
	return sum
}

func main() {
	var m, n int
	for {
		fmt.Printf("enter the m and n : ")
		_, err := fmt.Scanf("%d %d", &m, &n)
		if err != nil {
			fmt.Println("fmt.Scanf() returned an error")
			break
		}
		result := utils.Timeit(createMatAndMultiply)(m, n)
		fmt.Printf("result: %f \n\n", result)
	}
}
