package main

import "fmt"

//when pass slice, map, channel. it will change the value within main(){} scope since the passed value
//is the reference of the type rather than the value

//int, string, etc are not reference type, they are primitive type. while the slice, map, channel, etc are reference type
func main() {
	m := make([]string, 1, 50)
	fmt.Println(m)
	changeReference(m)
	fmt.Println(m)
}

func changeReference(z []string) {
	z[0] = "Todd"
	fmt.Println(z)
}
