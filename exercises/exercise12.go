package main

import "fmt"

// Escreva um programa que determine se um número é perfeito.
func main() {
	var number int

	fmt.Print("Digite um número: ")

	_, err := fmt.Scan(&number)
	if err != nil {
		panic(err)
	}

	var sum int
	for i := 1; i < number; i++ {
		if number%i == 0 {
			sum += i
		}
	}

	if sum == number {
		fmt.Print("Número perfeito")
	} else {
		fmt.Print("Número não perfeito")
	}
}
