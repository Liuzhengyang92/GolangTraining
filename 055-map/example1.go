package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println(v1)

	fmt.Println("length: ", len(m))

	delete(m, "k2")
	fmt.Println("map: ", m)

	_, ok := m["k2"]//due to the value at with key "k2" has been deleted, ok will return false
	fmt.Println("prs: ", ok)

}
