package main

import "fmt"

// Crie um programa que exiba os números pares de 1 a 50.
func main() {
	for i := 0; i < 50; i++ {
		number := i + 1

		if number%2 == 0 {
			fmt.Println(number)
		}
	}
}
