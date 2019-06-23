package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("../temp.txt")//create file named temp.txt. relative path is ../temp.txt
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	r := strings.NewReader("wassssup")

	io.Copy(f, r)
}
