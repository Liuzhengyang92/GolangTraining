package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	j := `[{"First": "James", "Last": "Bond", "Age": 32, "Sayings": ["Shaken", "KFC"]}]`
	fmt.Println(j)

	var people []person

	err := json.Unmarshal([]byte(j), &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, person := range people {
		fmt.Println("Person #", i)
		fmt.Println("\t", person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println(saying)
		}
	}
}
