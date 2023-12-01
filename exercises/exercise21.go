package main

import "fmt"

// Faça um programa que calcule o produto escalar de dois arrays.
func main() {
	var arr1 = array("Vetor A")
	var arr2 = array("Vetor B")

	var product int

	if len(arr1) != len(arr2) {
		product = -1
	}

	for i := 0; i < len(arr1); i++ {
		product += arr1[i] * arr2[i]
	}

	fmt.Printf("O produto escalar é %d", product)
}

func array(name string) (arr []int) {
	var length int

	fmt.Printf("%s - Digite o tamanho do vetor: ", name)

	_, err := fmt.Scan(&length)
	if err != nil {
		panic(err)
	}

	for i := 0; i < length; i++ {
		var number int

		fmt.Print("Digite o número: ")

		_, err := fmt.Scan(&number)
		if err != nil {
			panic(err)
		}

		arr = append(arr, number)
	}

	return
}
