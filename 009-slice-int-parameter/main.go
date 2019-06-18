package main

import "fmt"

func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is", x)
}

//...int  a slice of int
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position, ", i, " we are now adding, ", v, " to the total value")
	}
	fmt.Println("The total is", sum)
	return sum
}
