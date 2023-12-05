package main

import "fmt"

// Escreva um programa que realize a ordenação por contagem (counting sort) em um array.
func main() {
	var length int

	fmt.Print("Qual o tamanho da lista: ")
	_, err := fmt.Scan(&length)
	if err != nil {
		panic(err)
	}

	var numbers = []int{}
	for i := 0; i < length; i++ {
		var number int

		fmt.Printf("Digite o %dº número: ", (i + 1))

		_, err = fmt.Scan(&number)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number)
	}

	maxArr := max(numbers)

	countArray := make([]int, maxArr+1)

	for _, num := range numbers {
		countArray[num]++
	}

	sortedArray := []int{}
	for i := 0; i <= maxArr; i++ {
		for countArray[i] > 0 {
			sortedArray = append(sortedArray, i)
			countArray[i]--
		}
	}

	fmt.Print(numbers, sortedArray)
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
