package main

import "fmt"

func main() {
	greeting := make([]string, 3, 5)
	greeting[0] = "Good morning"
	greeting[1] = "Bonjour"
	greeting[2] = "beunos dias"

	greeting[3] = "zao an"
	//here is an error since the length of the slice is 3.
	//if we need to enlarge its length, we can use append: greeting = append(greeting, "zao an")

	fmt.Println(greeting)


	mySlice := []int{1, 2, 3, 4}
	otherSlice := []int{5, 6, 7}

	mySlice = append(mySlice, otherSlice...)
	fmt.Println(mySlice)

	newSlice := append(mySlice[2:], otherSlice[3:]...)
	fmt.Println(newSlice)
}
