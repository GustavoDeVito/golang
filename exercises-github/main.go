package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	var option int

	fmt.Print("Digite o exercicio que deseja executar: ")

	_, err := fmt.Scan(&option)
	if err != nil {
		panic(err)
	}

	switch option {
	case 1:
		exercise01()
		break

	case 2:
		exercise02()
		break

	case 3:
		exercise03()
		break

	case 4:
		exercise04()
		break

	case 5:
		exercise05()
		break

	default:
		fmt.Printf("Opção invalida!")
	}

	os.Exit(0)
}

func exercise01() {
	fmt.Print("Hello, world!")
}

func exercise02() {
	var numbers = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(numbers)
}

func exercise03() {
	var text string

	fmt.Print("Insira qualquer texto: ")

	_, err := fmt.Scan(&text)
	if err != nil {
		panic(err)
	}

	reverseText := reverse(text)

	fmt.Println(reverseText)
}

func exercise04() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, v := range lines {
		fmt.Println(v)
	}

	fmt.Printf("Total de %d linhas", len(lines))

	file.Close()
}

func exercise05() {
	file, err := os.Open("apache.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var clients = []string{}
	for _, v := range lines {
		data := strings.Split(v, " ")

		if !slices.Contains(clients, data[0]) {
			clients = append(clients, data[0])
		}
	}

	fmt.Printf("Total de clientes: %d", len(clients))

	file.Close()
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}

	return
}
