package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	fmt.Println(foo(aSlice...))
	foo()
}

func foo(numbers ...int) int {
	fmt.Println(numbers)
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return  sum
}