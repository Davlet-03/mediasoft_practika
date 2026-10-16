package main

import "testing"

func TestGenerateMatrix(t *testing.T) {
	matrix := generateMatrix()
	seen := make(map[int]bool)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			value := matrix[i][j]

			if value < 1 || value > rows*cols {
				t.Errorf("значение %d вне диапазона", value)
			}

			if seen[value] {
				t.Errorf("найден повтор числа %d", value)
			}

			seen[value] = true
		}
	}

	if len(seen) != rows*cols {
		t.Errorf("уникальных чисел %d, ожидали %d", len(seen), rows*cols)
	}
}
