package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(5)
	fmt.Println(runtime.GOMAXPROCS(0))
}
