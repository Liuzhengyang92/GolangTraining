package main

import "fmt"

func main() {
	age := 44
	fmt.Println(age)
	changeMe(&age)
	fmt.Println(age)
}

func changeMe(z *int) {
	//if the given value is z int, the value of z within the main(){} scope will not
	//be changed. If the given value is the pointer of z, the value will of z within scope main(){} will be
	//changed by *z=24
	fmt.Println(z)
	fmt.Println(*z)
	*z = 24
}
