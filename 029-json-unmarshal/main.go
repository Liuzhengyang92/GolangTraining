package main

import (
	"encoding/json"
	"fmt"
)

//Upper case is required to Unmarshal, if using lower case, the value will not be picked up through Unmarshal
//in this case, age does not convert to v.age successfully
type person struct {
	First string
	Last string
	age int
}

func main() {
	s := `[{"First":"Miss","Last":"Moneypenny","Age":27},{"First":"James","Last":"Bond","Age":32}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//people := []person{}

	var people []person

	err := json.Unmarshal(bs, &people)//give the converted value bs to the pointer &people

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("all of the data", people)

	for i, v := range people {
		fmt.Println("----PERSON NUMBER", i)
		fmt.Println(v.First, v.Last, v.age)
	}
}
