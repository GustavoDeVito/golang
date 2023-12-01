package main

import "fmt"

// Crie um programa que multiplique duas matrizes.
func main() {
	matrixA := inputMatrix("Matriz A")
	matrixB := inputMatrix("Matriz B")

	if len(matrixA[0]) == len(matrixB) {
		rows := len(matrixA)
		cols := len(matrixB[0])

		matrixC := make([][]int, rows)
		for i := range matrixC {
			matrixC[i] = make([]int, cols)
		}

		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				for k := 0; k < len(matrixA[0]); k++ {
					matrixC[i][j] += matrixA[i][k] * matrixB[k][j]
				}
			}
		}

		printMatrix(matrixA)
		println("-------------------------------------------------------")
		printMatrix(matrixB)
		println("-------------------------------------------------------")
		printMatrix(matrixC)
	} else {
		fmt.Print("Produto não existe")
	}
}

func inputMatrix(matrixName string) [][]int {
	var rows, cols int

	fmt.Printf("--%s--\n", matrixName)

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

	return matrix
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}
