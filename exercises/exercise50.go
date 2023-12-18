package main

import "fmt"

// Implemente um programa que calcule a sequência de Hofstadter Q.
func main() {
	var num int

	fmt.Print("Digite un número: ")

	_, err := fmt.Scan(&num)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Os primeiros %d termos da sequência de Hofstadter Q são:\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Q(%d) = %d\n", i, hofstadterQ(i))
	}
}

func hofstadterQ(n int) int {
	if n <= 1 {
		return 1
	}

	return hofstadterQ(n-hofstadterQ(n-1)) + hofstadterQ(n-hofstadterQ(n-2))
}
