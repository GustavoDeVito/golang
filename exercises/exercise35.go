package main

import (
	"fmt"
	"math"
)

// Crie um programa que verifique se um número é um quadrado perfeito.
func main() {
	var num int

	fmt.Print("Qual a medida de um dos lados do quadrado: ")

	_, err := fmt.Scan(&num)
	if err != nil {
		panic(err)
	}

	square := int(math.Sqrt(float64(num)))

	isPerfectSquare := square*square == num

	if isPerfectSquare {
		fmt.Print("É um quadrado perfeito")
	} else {
		fmt.Print("Não é um quadrado perfeito")
	}
}
