package main

import "fmt"

func main() {
	c := make(chan int, 1)//1 means only one buffer, if given 2, the following commands can be executed properly

	c <- 42
	c <- 43

	fmt.Println(<-c)
}

//this main function cannot be executed properly
//the channle only has one buffer. it has been given a value 42
//when trying to give it another value 43. so it waits to get the value off
//from the channel, but there is a call to set another value 43 to it.
//then there is a deadlock