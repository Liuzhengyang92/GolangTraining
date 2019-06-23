package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := generate(q)

	receivevalue(c, q)
	fmt.Println("about to exit")
}

func receivevalue(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}

	}
}

func generate(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i<100; i++ {
			c <- i
		}
		q <- 1
		close(c)
	}()
	return c

}
