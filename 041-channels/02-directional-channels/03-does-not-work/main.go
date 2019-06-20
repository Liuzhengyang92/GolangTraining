package main

import "fmt"

func main() {
	c := make(<-chan int, 2) //the channel c is a directional channel. it is a send-only channel
	fmt.Printf("%T\n", c)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("--------")
}
