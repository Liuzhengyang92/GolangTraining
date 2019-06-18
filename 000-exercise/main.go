package main

import "fmt"

//func main() {
//	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
//	for i, v := range x {
//		fmt.Println(i, "\t", v)
//	}
//	fmt.Printf("%T\n", x)
//	fmt.Println(x[1:6])
//}

//func main(){
//	x := make([]string, 50, 50)
//	x = []string{"a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a","a", }
//	fmt.Println(x)
//	fmt.Println(len(x))
//	fmt.Println(cap(x))
//
//	for i := 0; i < len(x); i++ {
//		fmt.Println(i, x[i])
//	}
//}

func main() {
	m := map[string][]string{
		`bond_james`:      {`Shaken, not strired`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice Cream`, `Sunsets`},
	}

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espinoage`}
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
