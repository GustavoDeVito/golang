package main

import "fmt"

func main() {
	var num int

	fmt.Print("Digite un número: ")

	_, err := fmt.Scan(&num)
	if err != nil {
		panic(err)
	}

	for num != 1 {
		if num%2 == 0 {
			num /= 2
		} else {
			num = (num * 3) + 1
		}

		fmt.Println(num)
	}
}
