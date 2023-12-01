package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Faça um programa que inverta um número inteiro.
func main() {
	var num int64

	fmt.Print("Digite um número: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		panic(err)
	}

	numStr := strconv.FormatInt(int64(num), 10)
	numbers := strings.Split(numStr, "")

	var reverse string

	for i := len(numbers); i > 0; i-- {
		reverse += numbers[i-1]
	}

	num, err = strconv.ParseInt(reverse, 10, 64)
	if err != nil {
		panic(err)
	}

	fmt.Println(num)
}
