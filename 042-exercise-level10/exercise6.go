package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c
	fmt.Println(v, ok)//42, true

	close(c)

	va, ok := <-c
	fmt.Println(va, ok)//0, false: since there is no value stored in the channel anymore.
}
