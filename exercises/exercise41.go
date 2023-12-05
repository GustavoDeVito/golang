package main

import "fmt"

// Faça um programa que verifique se um número é positivo, negativo ou zero.
func main() {
	var number int

	fmt.Print("Digite um número: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		panic(err)
	}

	if number > 0 {
		fmt.Printf("O número %d é positivo", number)
	} else if number < 0 {
		fmt.Printf("O número %d é negativo", number)
	} else {
		fmt.Printf("O número %d não é positivo e nem negativo", number)
	}
}
