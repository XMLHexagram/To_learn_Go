package main

import "fmt"

func swap(a, b *int) {
	fmt.Println(a, b, *a, *b)
	b, a = a, b
	fmt.Println(a, b, *a, *b)
}

func main() {
	x, y := 1, 2
	fmt.Println(&x, &y, x, y)
	swap(&x, &y)
	fmt.Println(&x, &y, x, y)
}
