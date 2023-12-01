package main

import "fmt"

// Faça um programa que verifique se uma matriz é simétrica.
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

	if isSymmetric(matrix) {
		fmt.Printf("Matriz simétrica")
	} else {
		fmt.Printf("Matriz antissimétrica")
	}
}

func isSymmetric(matrix [][]int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	if rows != cols {
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] != matrix[j][i] {
				return false
			}
		}
	}

	return true
}
