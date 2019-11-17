package main

import "fmt"

//
type T struct {
	a int
	b int
}

func (t T) getAndPrint(a int) (int, int) {
	return t.a, a
}

func (t T) get() int {
	return t.a
}

func (t *T) set(a int) {
	t.a = a
}

//

func main() {
	var temp T
	f := temp.set
	f(55)
	//temp.set(666)
	fmt.Printf("%p,%v,%d,%d\n",
		&temp, temp, temp.get(), temp.a)
	fmt.Println(temp.getAndPrint(66))
}
