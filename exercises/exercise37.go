package main

import "fmt"

// Faça um programa que conte quantas vezes um número aparece em um array.
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

	var number int

	fmt.Print("Digite um número para procurar na lista: ")

	_, err = fmt.Scan(&number)
	if err != nil {
		panic(err)
	}

	var count int
	for _, v := range numbers {
		if number == v {
			count++
		}
	}

	fmt.Printf("O número %d aparece %d vezes na lista", number, count)
}
