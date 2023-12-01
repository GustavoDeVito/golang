package main

import (
	"fmt"
	"math"
)

// Faça um programa que calcule a raiz quadrada de um número usando o método de Newton.
func main() {
	var number float64

	fmt.Print("Digite um número: ")

	_, err := fmt.Scan(&number)
	if err != nil {
		panic(err)
	}

	var y = 1.0
	for i := 0; i < 100; i++ {
		fy := y*y - number
		fPrimeY := 2 * y

		if math.Abs(fy) < 1e-10 {
			break
		}

		y = y - fy/fPrimeY
	}

	fmt.Printf("A raiz quadrada de %.2f é aproximadamente %.3f", number, y)
}
