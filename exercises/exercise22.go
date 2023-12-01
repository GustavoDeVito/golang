package main

import (
	"fmt"
	"reflect"
)

// Implemente um programa que verifique se dois arrays são iguais.
func main() {
	var arry1 = []int{}
	var arry2 = []int{}

	var length int

	fmt.Print("Digite a quantidade de números que deseja inserir: ")

	_, err := fmt.Scan(&length)
	if err != nil {
		panic(err)
	}

	for i := 0; i < length; i++ {
		var number int

		fmt.Printf("Vetor 1 - Insira o %d número: ", (i + 1))

		_, err = fmt.Scan(&number)
		if err != nil {
			panic(err)
		}

		arry1 = append(arry1, number)
	}

	for i := 0; i < length; i++ {
		var number int

		fmt.Printf("Vetor 2 - Insira o %d número: ", (i + 1))

		_, err = fmt.Scan(&number)
		if err != nil {
			panic(err)
		}

		arry2 = append(arry2, number)
	}

	if reflect.DeepEqual(arry1, arry2) {
		fmt.Print("Array Iguais")
	} else {
		fmt.Print("Array Diferentes")
	}
}
