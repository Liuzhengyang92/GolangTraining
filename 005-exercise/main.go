package main

import "fmt"

type person struct {
	first string
	last string
	favFlavours []string
}

func main() {
	p1 := person {
		first: "James",
		last: "Bond",
		favFlavours: []string{
			"chocolate",
			"martini",
		},
	}

	p2 := person {
		first: "Miss",
		last: "Moneypenny",
		favFlavours: []string{
			"strawberry",
			"vanilla",
		},
	}

	m := map[string]person{
		p1.last:p1,
		p2.last:p2,
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}

	//fmt.Println(p1.first)
	//fmt.Println(p1.last)
	//fmt.Println(p1.favFlavours)
	//for i, v := range p1.favFlavours {
	//	fmt.Println(i, v)
	//}
	//
	//fmt.Println(p2.first)
	//fmt.Println(p2.last)
	//fmt.Println(p2.favFlavours)
	//for i, v := range p2.favFlavours {
	//	fmt.Println(i, v)
	//}
}


