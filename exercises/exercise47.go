package main

import "fmt"

// Crie um programa que verifique se uma matriz é diagonal.
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

	if isDiag(matrix) {
		fmt.Printf("Matriz é diagonal")
	} else {
		fmt.Printf("Matriz não é diagonal")
	}
}

func isDiag(matrix [][]int) bool {
	rows := len(matrix)
	cols := len(matrix[0])

	if rows != cols {
		return false
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i != j && matrix[i][j] != 0 {
				return false
			}
		}
	}

	return true
}
