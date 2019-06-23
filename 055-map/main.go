package main

import "fmt"

func main() {
	m := make(map[string]string)

	m["Kobe"] = "Bryant"
	m["Lebron"] = "James"
	m["Michael"] = "Jordan"
	m["Anthony"] = "Davis"

	fmt.Println(m)

	m["Michael"] = "Jackson"

	fmt.Println(m)
}
