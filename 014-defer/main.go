package main

import "fmt"

func main() {
	//defer means to be excuted later. in this case, foo() will be executed at the end of main() function
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
