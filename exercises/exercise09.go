package main

import (
	"fmt"
	"strings"
)

// Escreva um programa que verifique se uma string é um palíndromo.
func main() {
	var text string

	fmt.Print("Digite um texto: ")

	_, err := fmt.Scan(&text)
	if err != nil {
		panic(err)
	}

	var reverse string

	words := strings.Split(text, "")
	for i := len(words); i > 0; i-- {
		reverse += words[i-1]
	}

	if text == reverse {
		fmt.Println("O texto é um palíndromo")
	} else {
		fmt.Println("O texto não é um palíndromo")
	}
}
