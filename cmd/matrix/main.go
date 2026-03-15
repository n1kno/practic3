package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fillUniqueMatrix(rows, cols, min, max int) [][]int {
	if max-min+1 < rows*cols {
		return nil
	}
	
	numbers := make([]int, max-min+1)
	for i := range numbers {
		numbers[i] = min + i
	}
	
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})
	
	matrix := make([][]int, rows)
	idx := 0
	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = numbers[idx]
			idx++
		}
	}
	
	return matrix
}

func main() {
	matrix := fillUniqueMatrix(3, 4, 1, 20)
	
	if matrix == nil {
		fmt.Println("Невозможно создать матрицу с уникальными значениями в заданном диапазоне")
		return
	}
	
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%2d ", matrix[i][j])
		}
		fmt.Println()
	}
}
