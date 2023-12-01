package main

import (
	"fmt"
)

// Implemente um programa que converta um número decimal para binário.
func main() {
	var number int

	fmt.Print("Digite um número: ")

	_, err := fmt.Scan(&number)
	if err != nil {
		panic(err)
	}

	fmt.Printf("O número %d em binário é %b", number, number)
}
