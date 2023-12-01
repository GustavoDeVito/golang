package main

import "fmt"

// Crie um programa que calcule a média de um conjunto de números.
func main() {
	var numbers = []int{}
	var length int

	fmt.Print("Quantos números deseja inserir: ")

	_, err := fmt.Scan(&length)
	if err != nil {
		panic(err)
	}

	for i := 0; i < length; i++ {
		var number int

		fmt.Printf("Digite o %dº número: ", (i + 1))

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
