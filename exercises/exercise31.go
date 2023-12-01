package main

import (
	"fmt"
	"strings"
)

// Crie um programa que conte quantas vezes uma substring aparece em uma string.
func main() {
	var text string

	fmt.Print("Digite um texto: ")

	_, err := fmt.Scan(&text)
	if err != nil {
		panic(err)
	}

	words := strings.Split(text, "")

	occurrences := make(map[string]int)
	for _, v := range words {
		_, ok := occurrences[v]

		if ok {
			occurrences[v] += 1
		} else {
			occurrences[v] = 1
		}
	}

	for key, value := range occurrences {
		fmt.Printf("%s: %v\n", key, value)
	}
}
