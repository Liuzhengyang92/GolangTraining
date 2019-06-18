package main

import "fmt"

func main() {
	x := sum([]int{2,3,4,5,6,7,8,9}...)
	fmt.Println("all numbers", x)

	s2 := even(sum, []int{2,3,4,5,6,7,8,9}...)
	fmt.Println("even numbers", s2)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v %2 == 0 {
			yi = append(yi, v)
		}
	}

	return f(yi...)
}