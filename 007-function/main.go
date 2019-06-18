package main

import "fmt"

func main() {
	bar()
	foo("James Bond")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

func bar() {
	fmt.Println("hello from function bar")
}

//everything in Go is pass by value
func foo(arg string) {
	fmt.Println("Hello", arg)
}

//the passed parameter is in the brackets with name and type, the return type is between brackets and curly brackets
func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, ` says, "Hello" `)
	b := false
	return a, b
}
