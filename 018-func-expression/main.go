package main

import "fmt"

func main() {
	fmt.Println("hello, playground")

	f := func() {
		fmt.Println("my first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("the year big brother started watching:", x)
	}
	g(1984)
}