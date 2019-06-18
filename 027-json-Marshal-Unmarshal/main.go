package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Animal struct {
	Name string
	Order string
}

type ColorGroup struct {
	ID int
	Name string
	Colors []string
}

func main() {
	var jsonBlob = []byte(`[
			{"Name": "Platypus", "Order": "Monotremata"},
			{"Name": "Quoll", "Order": "Dasyuromorphia"}
		]`)

	var animals []Animal

	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Printf("%+v\n", animals)


	group := ColorGroup{
		ID: 1,
		Name: "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(b)
}
