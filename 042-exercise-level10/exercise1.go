package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int, 1)
	fmt.Println("number of routine: ", runtime.NumGoroutine())
	go func() {
		fmt.Println("number of routines: ", runtime.NumGoroutine())
		c <- 42
		fmt.Println("give value 42 to the channel")
		c <- 43
		fmt.Println("give value 43 to the channel")
	}()

	fmt.Println("number of routines at the end", runtime.NumGoroutine())

	fmt.Println(<-c)//if <-c is not given, the go func() will not be executed.
	//here the c is blocked by <-. it is blocked and waits for the value of c to take out. when finish executing the
	//c <-  in go func(), the c will be released

	fmt.Println("print the first value of channel")
	fmt.Println(<-c)
	fmt.Println("print the second value of channel")
}
