package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 3, 6, 1, 4}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr.no"}

	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)

	fmt.Println("----------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
