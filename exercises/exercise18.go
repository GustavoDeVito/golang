package main

import (
	"fmt"
	"reflect"
	"sort"
)

// Implemente um programa que verifique se um array está ordenado.
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

	numbersSorted := make([]int, len(numbers))

	copy(numbersSorted, numbers)

	sort.Ints(numbersSorted)

	if reflect.DeepEqual(numbers, numbersSorted) {
		fmt.Print("Array ordenado")
	} else {
		fmt.Print("Array não ordenado")
	}
}
