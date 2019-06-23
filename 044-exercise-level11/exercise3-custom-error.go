package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func(ce customErr) Error() string {
	return fmt.Sprintf("here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more coffee",
	}

	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran - ", e, "\n", e.(customErr).info) //if given e.info there will be an error: error does not have a field info
	//if we want to access the info of error. we need to use insertion. e.(customErr).info()
}
