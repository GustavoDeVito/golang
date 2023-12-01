package main

import (
	"fmt"
	"slices"
)

// Implemente um programa que calcule o máximo divisor comum (MDC) de dois números.
func main() {
	var num1 int
	var num2 int

	fmt.Print("Digite o primeiro número: ")
	_, err := fmt.Scan(&num1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Digite o segundo número: ")
	_, err = fmt.Scan(&num2)
	if err != nil {
		panic(err)
	}

	var dividersNum1 = []int{}
	var dividersNum2 = []int{}

	for i := 1; i <= num1; i++ {
		if num1%i == 0 {
			dividersNum1 = append(dividersNum1, i)
		}
	}

	for i := 1; i <= num2; i++ {
		if num2%i == 0 {
			dividersNum2 = append(dividersNum2, i)
		}
	}

	var commonDivisors = []int{}
	for _, v := range dividersNum1 {
		if slices.Contains(dividersNum2, v) {
			commonDivisors = append(commonDivisors, v)
		}
	}

	fmt.Printf("MDC: %d", max(commonDivisors))
}

func max(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return max
}
