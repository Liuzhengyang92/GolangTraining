package main

import "fmt"

func main() {
	fmt.Println( half(6))
	//h, even := half(5)
	//fmt.Println(h, even)
}

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}
