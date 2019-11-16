package main

import "fmt"

type IntSlice []int

func (s IntSlice) sum() int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

func sumWithoutMethod(s []int) int {
	sum := 0
	for _, i := range s {
		sum += i
	}
	return sum
}

func main() {
	var temp []int
	temp = append(temp, 1)
	temp = append(temp, 2)
	temp = append(temp, 3)
	fmt.Printf("%d\n", sumWithoutMethod(temp))
	var a IntSlice = temp
	fmt.Printf("%d\n", a.sum())
}
