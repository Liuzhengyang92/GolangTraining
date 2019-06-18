package main

import "fmt"

func main() {
		xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
		x := sum(xi...)//a slice of int
		//x := sum(2, 3, 4, 5, 6, 7, 8, 9)
		fmt.Println("The total is ", x)
}

//the intake of the sum method is unlimited int values
func sum(x ...int) int {
	fmt.Println(x)
	fmt.Println("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position ", i, "we are now adding ", v, "to the total")
	}

	return sum
}
