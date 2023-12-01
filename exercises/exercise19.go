package main

import (
	"fmt"
	"slices"
	"strings"
)

// Crie um programa que conte a quantidade de vogais em uma string.
func main() {
	var vowels = []string{"a", "e", "i", "o", "u"}
	var text string

	fmt.Print("Digite um texto: ")

	_, err := fmt.Scan(&text)
	if err != nil {
		panic(err)
	}

	words := strings.Split(text, "")

	var count int
	for _, v := range words {
		if slices.Contains(vowels, v) {
			count++
		}
	}

	fmt.Printf("Vogais: %d", count)
}
