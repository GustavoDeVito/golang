package main

import "fmt"

// Escreva um programa que encontre o número mais frequente em uma matriz.
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

	occurrences := make(map[int]int)
	for _, v := range numbers {
		_, ok := occurrences[v]

		if ok {
			occurrences[v] += 1
		} else {
			occurrences[v] = 1
		}
	}

	var max int
	for k, v := range occurrences {
		if v > occurrences[max] {
			max = k
		}
	}

	fmt.Printf("Número com mais ocorrência: %d", max)
}
