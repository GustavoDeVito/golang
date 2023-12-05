package main

import (
	"fmt"
	"math"
)

// Faça um programa que verifique se um número é primo de Mersenne.
func main() {
	var number int

	fmt.Print("Digite um número: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		panic(err)
	}

	if isMersennePrime(number) {
		fmt.Printf("O número %d é primo de Mersenne", number)
	} else {
		fmt.Printf("O número %d não é primo de Mersenne", number)
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isMersennePrime(p int) bool {
	if !isPrime(p) {
		return false
	}
	mersenne := int(math.Pow(2, float64(p))) - 1
	return isPrime(mersenne)
}
