package main

import "fmt"

// Faça um programa que determine se um número é primo.
func main() {
	var number int

	fmt.Print("Digite um número: ")

	_, err := fmt.Scan(&number)
	if err != nil {
		panic(err)
	}

	if number < 2 || number%2 == 0 || number%3 == 0 {
		fmt.Printf("Número %d não é primo", number)
	} else {
		fmt.Printf("Número %d é primo", number)
	}
}
