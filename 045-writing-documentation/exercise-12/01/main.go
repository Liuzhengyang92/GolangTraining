package main

import (
	"GolangTraining/045-writing-documentation/exercise-12/01/dog"
	"fmt"
)

type canine struct {
	name string
	age int
}
func main() {
	fido := canine{
		name: "Fido",
		age: dog.Years(10),
	}

	fmt.Println(fido)
}
