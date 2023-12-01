package main

import "fmt"

// Escreva um programa que exiba a sequência de Fibonacci com os primeiros 20 termos.
func main() {
	var sequence = []int{}

	for i := 1; i <= 18; i++ {
		var sum int

		if i-2 < 0 {
			sum = 1 + 1

			sequence = append(sequence, 1)
			sequence = append(sequence, 1)
		} else {
			n1 := i + 1 - 1
			n2 := i + 1 - 2

			sum = sequence[n1] + sequence[n2]
		}

		sequence = append(sequence, sum)
	}

	fmt.Print(len(sequence), " - ", sequence)
}
