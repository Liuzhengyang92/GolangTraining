package main

import "fmt"

//pointer is the position of the data is stored, it is a constant whatever the data is changed to
func main() {
	y := 0
	boo(&y)
	fmt.Println(y)

	pointer := foo(y)
	fmt.Println(pointer)
}

func foo(x int) *int {
	fmt.Println(x)
	x = 43
	fmt.Println(x)
	return &x
}

func boo(x *int) {
	fmt.Println(x)
	*x = 43
	fmt.Println(x)
}
