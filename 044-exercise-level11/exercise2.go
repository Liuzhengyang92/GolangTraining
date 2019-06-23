package main

import (
	"encoding/json"
	"fmt"
)

type persons struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := persons{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"this", "is", "my", "house"},
	}

	bs, err := toJSON(p1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("There was an error in toJSON: %v", err)
	}

	return bs, nil
}
