package main

import (
	"fmt"
	"sort"
)

// Implemente um programa que calcule a mediana de um conjunto de números.
func main() {
	var length int

	fmt.Print("Qual o tamanho da lista: ")
	_, err := fmt.Scan(&length)
	if err != nil {
		panic(err)
	}

	var numbers = []int{}
	for i := 0; i < length; i++ {
		var number int

		fmt.Printf("Digite o %dº número: ", (i + 1))

		_, err = fmt.Scan(&number)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number)
	}

	sort.Ints(numbers)

	if length%2 == 0 {
		num1 := numbers[length/2-1]
		num2 := numbers[length/2]

		calc := (num1 + num2) / 2

		fmt.Printf("A mediana dos valores %v é %d", numbers, calc)
	} else {
		fmt.Printf("A mediana dos valores %v é %d", numbers, numbers[int(length/2)])
	}
}
