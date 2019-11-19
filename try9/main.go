package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func randIntBuilder(done chan struct{}) chan int {
	temp := make(chan int)
	go func() {
	tempL:
		for {
			select {
			case temp <- rand.Int():
			case <-done:
				break tempL
			}
		}
		fmt.Printf("end1\n")
		close(temp)
		fmt.Printf("end2\n")
	}()
	return temp
}

func main() {
	done := make(chan struct{})
	ch := randIntBuilder(done)

	fmt.Println("NumGoroutine", runtime.NumGoroutine(), " ", 0)
	fmt.Println(<-ch, " ", 1)
	fmt.Println(<-ch, " ", 2)

	close(done)
	fmt.Println(<-ch, " ", 3)
	//time.Sleep(3 * time.Second)
	fmt.Println(<-ch, " ", 4)
	fmt.Println(<-ch, " ", "test")

	fmt.Println("NumGoroutine", runtime.NumGoroutine(), " ", 5)
}
