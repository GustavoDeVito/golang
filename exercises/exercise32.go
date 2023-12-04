package main

import (
	"fmt"
	"slices"
)

// Escreva um programa que realize a busca binária em um array ordenado.
func main() {
	var list = []int{}

	var length int

	fmt.Print("Digite o tamanho da lista: ")

	_, err := fmt.Scan(&length)
	if err != nil {
		panic(err)
	}

	for i := 0; i < length; i++ {
		var number int

		fmt.Printf("Digite o %d número da lista: ", (i + 1))

		_, err := fmt.Scan(&number)
		if err != nil {
			panic(err)
		}

		list = append(list, number)
	}

	var element int

	fmt.Print("Qual elemento deseja buscar dentro da lista: ")

	_, err = fmt.Scan(&element)
	if err != nil {
		panic(err)
	}

	newList := removeDuplicates(list)

	slices.Sort(newList)

	result := binarySearch(newList, element)

	if result != -1 {
		fmt.Printf("O elemento está na posição %d", result)
	} else {
		fmt.Printf("O elemento %d não está na lista", element)
	}
}

func removeDuplicates(list []int) []int {
	var newList = []int{}

	for _, v := range list {
		if !slices.Contains(newList, v) {
			newList = append(newList, v)
		}
	}

	return newList
}

func binarySearch(list []int, element int) int {
	left, right := 0, len(list)-1

	for left <= right {
		mid := (left + right) / 2

		if list[mid] == element {
			return mid
		} else if list[mid] < element {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
