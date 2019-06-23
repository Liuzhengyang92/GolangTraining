package main

import "fmt"

func main() {
	func() {
		fmt.Println("this is the output of an anonymous function")
	}()
}
