package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First string
	Last string
	Sayings []string
}

func main() {
	p1 := person {
		First: "James",
		Last: "Bond",
		Sayings:[]string{"This", "is", "my", "house"},
	}

	bs, err := json.Marshal(p1)

	if err != nil {
		log.Fatalln("The JSON does not marshal, here is an error: ", err)
	}

	fmt.Println(string(bs))
}
