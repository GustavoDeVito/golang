package main

import "fmt"

// Calcule a soma dos números de 1 a 100.
func main() {
	n := 100
	a1 := 1
	aN := 100

	sum := n / 2 * (a1 + aN)

	fmt.Printf("A soma é %d", sum)
}
