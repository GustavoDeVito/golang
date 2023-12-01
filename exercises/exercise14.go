package main

import "fmt"

// Implemente um programa que converta uma temperatura em Celsius para Fahrenheit.
func main() {
	var celsius float64

	fmt.Print("Digite a temperatura em Celsius: ")

	_, err := fmt.Scan(&celsius)
	if err != nil {
		panic(err)
	}

	fahrenheit := (celsius * 9 / 5) + 32

	fmt.Printf("Temperatura em Fahrenheit: %.1f", fahrenheit)
}
