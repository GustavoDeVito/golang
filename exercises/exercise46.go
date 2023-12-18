package main

import "fmt"

// Implemente um programa que faça a rotação de uma matriz 90 graus no sentido horário.
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

	rotateMatrix := make([][]int, rows)
	for i := range rotateMatrix {
		rotateMatrix[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotateMatrix[j][rows-1-i] = matrix[i][j]
		}
	}

	fmt.Println("Matriz Original:")
	printMatrix(matrix)

	fmt.Println("Matriz Rotacionada 90 graus no Sentido Horário:")
	printMatrix(rotateMatrix)

}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}

	fmt.Print("\n")
}
