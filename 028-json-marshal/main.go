package main

import (
	"encoding/json"
	"fmt"
)

//the struct field must start with a upper case, or the json.Marshal cannot pick it up. when giving lower case, it
//won't return value after Marshal
//This can be used to select the values that we want to convert from struct and json
type person struct {
	First string
	Last string
	age int
}

func main() {
	p1 := person{
		First: "Miss",
		Last: "Moneypenny",
		age: 27,
	}

	p2 := person{
		First: "James",
		Last: "Bond",
		age: 32,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
