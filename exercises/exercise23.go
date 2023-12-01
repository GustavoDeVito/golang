package main

import "fmt"

// Crie um programa que calcule a média de uma matriz.
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

	var sum int
	for _, v := range numbers {
		sum += v
	}

	fmt.Printf("A média é %d", (sum / length))
}
