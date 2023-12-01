package main

import (
	"fmt"
	"sort"
)

// Faça um programa que encontre o maior e o menor elemento em um array.
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

	min, max := numbers[0], numbers[len(numbers)-1]

	fmt.Printf("O menor valor é %d e o maior é %d", min, max)
}
