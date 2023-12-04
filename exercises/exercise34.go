package main

import "fmt"

// Implemente um programa que faça a rotação de um array para a direita.
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

		fmt.Printf("Digite o %dº número da lista: ", (i + 1))

		_, err := fmt.Scan(&number)
		if err != nil {
			panic(err)
		}

		list = append(list, number)
	}

	var rotations int

	fmt.Print("Digite o número de rotações que deseja: ")

	_, err = fmt.Scan(&rotations)
	if err != nil {
		panic(err)
	}

	var newList = []int{}
	for i := len(list) - rotations; i < len(list); i++ {
		newList = append(newList, list[i])
	}

	for i := 0; i < len(list)-rotations; i++ {
		newList = append(newList, list[i])
	}

	fmt.Print(newList)
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
