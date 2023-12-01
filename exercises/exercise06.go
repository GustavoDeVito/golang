package main

import "fmt"

// Crie um programa que calcule o fatorial de um número.
func main() {
	var number int

	fmt.Print("Digite um número: ")

	_, err := fmt.Scan(&number)
	if err != nil {
		panic(err)
	}

	if number == 0 {
		number = 1
	}

	factorial := number
	for i := number - 1; i > 0; i-- {
		factorial *= i
	}

	fmt.Println(factorial)
}
