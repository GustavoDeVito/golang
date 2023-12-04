package main

import "fmt"

// Implemente um programa que realize a multiplicação de dois números sem usar o operador de multiplicação.
func main() {
	var num1, num2 int

	fmt.Print("Digite o primeiro valor: ")

	_, err := fmt.Scan(&num1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Digite o segundo valor: ")

	_, err = fmt.Scan(&num2)
	if err != nil {
		panic(err)
	}

	var result int
	for i := 1; i <= num2; i++ {
		result += num1
	}

	fmt.Printf("Resultado: %d", result)
}
