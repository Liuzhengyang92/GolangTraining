package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello everyone, this is the first class of my go experience", 42, true)
	fmt.Println(n)
	fmt.Println(err)

	foo()

	fmt.Println("something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()

	y := 100 + 24
	fmt.Println(y)

}

func foo() {
	fmt.Println("I'm a fool")
}

func bar() {
	n, _ := fmt.Println("and then we exited", 42, true)
	fmt.Println(n)
}

// control flow:
// (1) sequence
// (2) loop; iterative
// (3) conditional
