package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age int
}

type ByAge []person

type ByName []person

func (bn ByName) Len() int {
	return len(bn)
}

func (bn ByName) Swap(i, j int) {
	bn[i], bn[j] = bn[j], bn[i]
}

func (bn ByName) Less(i, j int) bool {
	return bn[i].Age < bn[j].Age
}

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func main() {
	p1 := person{"James", 32}
	p2 := person{"Moneypenny", 27}
	p3 := person{"Q", 64}
	p4 := person{"M", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)

	//Sort sorts data. It makes one call to data.Len to determine n, and O(n*log(n))
	//calls to data.Less and data.Swap.
	//sort.Sort(ByAge(people))
	//fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)
}
