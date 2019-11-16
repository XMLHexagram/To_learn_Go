package main

import "fmt"

type Map map[string]string

func (m Map) Print() {
	fmt.Println(m)
}

type iMap Map

func (m iMap) Print() {
	fmt.Println(m)
}

type slice_ []int

func (s slice_) Print() {
	fmt.Println(s)
}

func main() {
	mp := make(map[string]string, 10)
	mp["hi"] = "tata"
	mp["try"] = "try"

	var ma Map = mp

	ma.Print()
	//var im iMap = ma
	var im iMap = (iMap)(ma)
	fmt.Printf("im\n")
	im.Print()

	var i interface {
		Print()
	} = ma

	i.Print()

	s1 := []int{1, 2, 3, 4, 5}
	var s2 slice_
	s2 = s1
	s2.Print()
}
