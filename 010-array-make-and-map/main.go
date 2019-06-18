package main

import "fmt"

func main() {
	//x := make([]int, 10, 100)
	//fmt.Println(x)
	//fmt.Println(len(x))
	//fmt.Println(cap(x))
	//
	//jb := []string{"James", "Bond", "Chocolate", "Vanilla"}
	//fmt.Println(jb)
	//
	//mp := []string{"Miss", "moneypenny", "Strawberry", "Hazelnut"}
	//fmt.Println(mp)
	//
	//xp := [][]string{jb, mp}
	//fmt.Println(xp)


	m := map[string]int{ //following map the first attri is the key type, the second is the value type
		"James": 32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"]) // if we give a key which is not existed, the value will be 0

	v, ok := m["Barnabas"] // value is to return the value for key "Barnabas", ok is to return whether the value with
							// key exists
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Miss Moneypenny"]; ok { // v, ok is to see whether the value with key "Miss Moneypenny" exists
											//then if ok is true, print the following line
		fmt.Println("This is the if print", v)
	}
}
