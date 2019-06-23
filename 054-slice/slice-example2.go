package main

import "fmt"

func main() {
	greeting := []string{
		"Good morning",
		"Bonjour",
		"Dias",
		"Bongionor",
		"Ohayo",
		"Selamat pagi",
		"Gutten morgen",
	}

	for i, currentEntry := range greeting{
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}

	fmt.Print("[1:2]")
	fmt.Println(greeting[1:2])
	fmt.Print("[:2]")
	fmt.Println(greeting[:2])
	fmt.Print("[:]")
	fmt.Println(greeting[:])//print everything in this slice
}
