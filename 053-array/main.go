package main

import "fmt"

//array is not dynamic because the size is a part of this type. while slice is dynamic since there is not size given
func main() {
	var x [58] string //given nunber in square brackets, this is an arrya, if not given, it is a slice
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 65; i < 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}
