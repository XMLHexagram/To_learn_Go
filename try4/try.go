package main

import "fmt"

type INT int

func (input INT) max(b INT) INT {
	if input >= b {
		return input
	} else {
		return b
	}
}
func (input INT) get() INT {
	return input
}

func (input *INT) set(a INT) {
	*input = a
}

func (input INT) print() {
	fmt.Printf("value = %d\n", input)
}

func main() {
	var temp INT
	var a INT
	temp.set(66)
	a = 99
	temp += a
	fmt.Printf("max = %d\n", temp.max(888))
	INT.print(temp)
	(*INT).print(&temp)
}
