package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}

//./exercise-06-os-arch will execute the code executable code exercise-06-os-arch
//rm exercise-06-os-arch will delete the executable app exercise-06-os-arch