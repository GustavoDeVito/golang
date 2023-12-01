package main

import (
	"fmt"
	"math"
)

// Faça um programa que calcule o cosseno de um ângulo usando a série de Taylor.
func main() {
	var angle float64
	var terms int

	fmt.Print("Digite o ângulo em graus: ")

	_, err := fmt.Scan(&angle)
	if err != nil {
		panic(err)
	}

	fmt.Print("Digite o número de termos na série de Taylor: ")

	_, err = fmt.Scan(&terms)
	if err != nil {
		panic(err)
	}

	angle = angle * (math.Pi / 180)
	result := 1.0

	for i := 1; i <= terms; i++ {
		term := math.Pow(angle, float64(2*i)) / float64(factorial(2*i))

		if i%2 == 0 {
			result += term
		} else {
			result -= term
		}
	}

	fmt.Printf("O cosseno de %f graus é aproximadamente: %f\n", angle, result)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
