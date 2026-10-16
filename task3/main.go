package main

import (
	"fmt"
	"math/rand"
)

const rows = 3
const cols = 4

func generateMatrix() [rows][cols]int {
	var matrix [rows][cols]int

	numbers := rand.Perm(rows * cols)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = numbers[i*cols+j] + 1
		}
	}

	return matrix
}

func main() {
	matrix := generateMatrix()

	fmt.Println("Двумерный массив:")

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Printf("%3d", matrix[i][j])
		}
		fmt.Println()
	}
}
