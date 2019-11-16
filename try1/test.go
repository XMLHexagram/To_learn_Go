package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	a := struct {
		name string
		age int
	}{"Alice",17}
	fmt.Printf("%T\n",a)
	fmt.Printf("%v\n",a)

	b := Person{"Ariel",17}
	fmt.Printf("%T\n",b)
	fmt.Printf("%v\n",b)
}
