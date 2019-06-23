package main

import (
	"fmt"
)

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	//the order of output is "hello world"; the reason
	//is that the scope of the defer, the first defer
	//applies for the whole scope of main, while the second
	//one applies within the first defer
	defer world()
	defer hello()
}
