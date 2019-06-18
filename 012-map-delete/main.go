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

	delete(m, "Ian Fleming")
	fmt.Println(m)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}

	fmt.Println(m)

	m["Miss"] = 18

	fmt.Println(m)
}
