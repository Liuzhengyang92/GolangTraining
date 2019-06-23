package main

import "fmt"

func main() {
	greeting := map[int]string{
		0: "Good morning",
		1: "Bonjour",
		2: "zao",
		3: "dias",
	}

	fmt.Println(greeting)

	if value, ok := greeting[2]; ok {
		fmt.Println(value)
		delete(greeting, 2)
	} else {
		fmt.Println("That value does not exist")
	}

	fmt.Println(greeting)
}
