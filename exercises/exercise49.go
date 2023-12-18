package main

import "fmt"

// Faça um programa que determine se um número é um quadrado mágico.
func main() {
	var rows, cols int

	fmt.Print("Digite a quantidade de linhas: ")

	_, err := fmt.Scan(&rows)
	if err != nil {
		panic(err)
	}

	fmt.Print("Digite a quantidade de colunas: ")

	_, err = fmt.Scan(&cols)
	if err != nil {
		panic(err)
	}

	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			var number int

			fmt.Printf("Digite um número para a linha %d e coluna %d: ", i+1, j+1)

			_, err := fmt.Scan(&number)
			if err != nil {
				panic(err)
			}

			matrix[i][j] = number
		}
	}

	if isSquareMagic(matrix) {
		fmt.Println("É um quadrado mágico!")
	} else {
		fmt.Println("Não é um quadrado mágico.")
	}
}

func isSquareMagic(square [][]int) bool {
	n := len(square)

	for _, row := range square {
		if len(row) != n {
			return false
		}
	}

	targetSum := sum(square[0])

	for _, row := range square {
		if sum(row) != targetSum {
			return false
		}
	}

	for j := 0; j < n; j++ {
		colSum := 0

		for i := 0; i < n; i++ {
			colSum += square[i][j]
		}

		if colSum != targetSum {
			return false
		}
	}

	diagonalSum := 0
	for i := 0; i < n; i++ {
		diagonalSum += square[i][i]
	}

	if diagonalSum != targetSum {
		return false
	}

	secondaryDiagonalSum := 0
	for i := 0; i < n; i++ {
		secondaryDiagonalSum += square[i][n-i-1]
	}

	if secondaryDiagonalSum != targetSum {
		return false
	}

	return true
}

func sum(arr []int) int {
	result := 0
	for _, num := range arr {
		result += num
	}
	return result
}
