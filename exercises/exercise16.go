package main

import (
	"fmt"
	"sort"
)

// Escreva um programa que ordene um array de números inteiros.
func main() {
	var numbers = []int{}
	var length int

	fmt.Print("Digite a quantidade de números que deseja inserir: ")

	_, err := fmt.Scan(&length)
	if err != nil {
		panic(err)
	}

	for i := 0; i < length; i++ {
		var number int

		fmt.Printf("Insira o %d número: ", (i + 1))

		_, err = fmt.Scan(&number)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number)
	}

	sort.Ints(numbers)

	fmt.Print(numbers)
}
