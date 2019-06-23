package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning",
		1: "Bonjour",
		2: "Zaoan",
		3: "Buerus dias",
	}

	for key, val := range myGreeting{
		fmt.Println(key, "-", val)
	}
}
