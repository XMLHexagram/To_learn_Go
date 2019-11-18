package main

import (
	"fmt"
	"runtime"
)

func sum10000(i chan struct{}) {
	var sum int = 0
	for i := 0; i <= 10000; i++ {
		sum += i
	}
	fmt.Println(sum)
	i <- struct{}{}
}

func main() {
	C := make(chan struct{})
	fmt.Println(runtime.GOMAXPROCS(0))
	go sum10000(C)

	fmt.Printf("test go\n")
	<-C
	//time.Sleep(5 * time.Second)
}
