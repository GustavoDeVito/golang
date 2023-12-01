package main

import (
	"fmt"
	"slices"
)

// Escreva um programa que remova elementos duplicados de um array.
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

	var unique = []int{}
	for _, v := range numbers {
		if !slices.Contains(unique, v) {
			unique = append(unique, v)
		}
	}

	fmt.Print(unique)
}
