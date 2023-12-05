package main

import (
	"fmt"
	"math"
)

// Faça um programa que calcule o valor do π usando a série de Leibniz.
func main() {
	iterations := 100000

	piApprox := calculatePi(iterations)

	piReal := math.Pi

	fmt.Printf("Valor aproximado de π com %d iterações: %f\n", iterations, piApprox)
	fmt.Printf("Valor real de π: %f\n", piReal)
}

func calculatePi(iterations int) float64 {
	piApproximation := 0.0
	sign := 1.0

	for i := 0; i < 100000; i++ {
		term := 1.0 / (2.0*float64(i) + 1.0)
		piApproximation += sign * term
		sign *= -1.0
	}

	return 4.0 * piApproximation
}
