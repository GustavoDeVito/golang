package main

import "fmt"

// Crie um programa que gere um triângulo de Pascal com um número de linhas dado.
func main() {
	var rows int

	fmt.Print("Digite a quantidade de linhas: ")

	_, err := fmt.Scan(&rows)
	if err != nil {
		panic(err)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(calculate(i, j), " ")
		}

		fmt.Println()
	}
}

func calculate(n, k int) int {
	if k == 0 || k == n {
		return 1
	}

	return calculate(n-1, k-1) + calculate(n-1, k)
}
