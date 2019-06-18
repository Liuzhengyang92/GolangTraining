package main

import "fmt"

func main() {

	m := map[string]int{ //following map the first attri is the key type, the second is the value type
		"James": 32,
		"Miss Moneypenny": 27,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "James")
	fmt.Println(m)
}