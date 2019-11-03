package main

import (
	"fmt"
)

func main() {
	var score int
	fmt.Printf("Please input your score:")
	fmt.Scanln(&score)

	if score >= 90 && score <= 100 {
		fmt.Printf("You get an A")
	} else if score < 90 && score >= 80 {
		fmt.Printf("You get a B")
	} else if score < 80 && score >= 70 {
		fmt.Printf("You get a C")
	} else if score < 70 && score >= 60 {
		fmt.Printf("You get a D")
	} else if score < 60 && score >= 0 {
		fmt.Printf("Sorry,you get an E")
	} else {
		fmt.Printf("You input wrong number")
	}
}
