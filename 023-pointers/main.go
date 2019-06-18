package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) //give you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b) // give you the value stored at an address when you have the address
	fmt.Println(*&a)
}
