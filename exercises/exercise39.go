package main

import (
	"fmt"
)

// Crie um programa que calcule o máximo divisor comum (MDC) de um conjunto de números.
func main() {
	var length int

	fmt.Print("Quantos número deseja inserir: ")
	_, err := fmt.Scan(&length)
	if err != nil {
		panic(err)
	}

	var numbers = []int{}
	for i := 0; i < length; i++ {
		var number int

		fmt.Printf("Digite o %dº número: ", (i + 1))

		_, err = fmt.Scan(&number)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number)
	}

	if len(numbers) < 2 {
		fmt.Println("Pelo menos dois números são necessários para calcular o MDC.")
	} else {
		mdc := calculateMDC(numbers)
		fmt.Printf("O MDC de %v é: %d", numbers, mdc)
	}
}

func max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return max
}

// Função para calcular o MDC de dois números usando o algoritmo de Euclides
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func calculateMDC(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = gcd(result, numbers[i])
	}

	return result
}
