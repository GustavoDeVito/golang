package main

import "fmt"

// Escreva um programa que determine se um número é uma potência de 2.
func main() {
	var number int

	fmt.Print("Digite um número: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		panic(err)
	}

	if isPowerOfTwo(number) {
		fmt.Printf("O número %d é uma potência de 2", number)
	} else {
		fmt.Printf("O número %d não é uma potência de 2", number)
	}
}

func isPowerOfTwo(num int) bool {
	return num > 0 && (num&(num-1)) == 0
}
