package main

import "fmt"

func main() {
	elements := map[string]map[string]string {
		"He": map[string]string{
			"Name": "Kevin",
			"State": "Marries",
		},
		"She": map[string]string{
			"Name": "Tom",
			"State": "Single",
		},
	}

	fmt.Println(elements)
	fmt.Println(elements["He"])
	fmt.Println(elements["She"])
}
