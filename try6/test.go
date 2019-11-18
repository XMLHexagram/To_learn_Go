package main

import (
	"fmt"
	"net/http"
)

func add(a, b int) int {
	return a + b
}

type ADD func(int, int) int

func testAdd(c ADD, a, b int) int {
	return c(a, b)
}
func main() {
	var temp ADD = add
	fmt.Println(temp(10, 20))
	fmt.Println(testAdd(temp, 30, 30))
	http.Handler()
}
