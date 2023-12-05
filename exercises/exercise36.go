package main

import "fmt"

// Escreva um programa que calcule a combinação de n elementos tomados de k em k.
func main() {
	var numN, numK int

	fmt.Print("Digite o valor de N: ")
	_, err := fmt.Scan(&numN)
	if err != nil {
		panic(err)
	}

	fmt.Print("Digite o valor de K: ")
	_, err = fmt.Scan(&numK)
	if err != nil {
		panic(err)
	}

	result := combinations(numN, numK)

	fmt.Printf("Combinação de %d elementos tomados de %d em %d: %d", numN, numK, numK, result)
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}

	result := 1

	for i := 1; i <= x; i++ {
		result *= i
	}

	return result
}

func combinations(n, k int) int {
	if n < k {
		return 0
	}

	return factorial(n) / (factorial(k) * factorial(n-k))
}
