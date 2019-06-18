package main

import "fmt"

type person struct {
	first string
}

func main() {
	p1 := person{
		first: "James Bond",
	}

	fmt.Println(p1)
	changeMe(&p1) //change the value of p1 from the global scope
	change(p1) //only change the value of p1 within the scope of function change
	fmt.Println(p1)
}

func changeMe(p *person) {
	p.first = "Miss Moneypenny"
}

func change(p person) {
	p.first = "Miss Moneypenny"
	fmt.Println(p)
}
