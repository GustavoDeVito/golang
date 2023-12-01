package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Implemente um programa que determine se um número é Armstrong.
func main() {
	var number float64

	fmt.Print("Digite um número: ")

	_, err := fmt.Scan(&number)
	if err != nil {
		panic(err)
	}

	numberStr := strings.Split((strconv.FormatInt(int64(number), 10)), "")

	var sum float64
	for _, v := range numberStr {
		num, _ := strconv.ParseFloat(v, 64)

		sum += math.Pow(num, float64(len(numberStr)))
	}

	if number == sum {
		fmt.Print("O número é Armstrong")
	} else {
		fmt.Print("O número não é Armstrong")
	}
}
