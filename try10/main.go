package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func generateA(done chan struct{}) chan int {
	ch := make(chan int)
	go func(chan int) {
	block:
		for true {
			select {
			case ch <- rand.Int():
			case <-done:
				break block
			}
		}
	}(ch)
	return ch
}

func generateB(done chan struct{}) chan int {
	ch := make(chan int)
	go func(chan int) {
	block:
		for true {
			select {
			case ch <- rand.Int():
			case <-done:
				break block
			}
		}
	}(ch)
	return ch
}

func mainGenerater(done chan struct{}) chan int {
	sendDone := make(chan struct{})
	ch := make(chan int)
	go func() {
	block:
		for {
			select {
			case ch <- <-generateA(sendDone):
			case ch <- <-generateB(sendDone):
			case <-done:
				//fmt.Printf("test%v\n", sendDone)
				sendDone <- struct{}{}
				//fmt.Printf("test%v\n", sendDone)
				sendDone <- struct{}{}
				//fmt.Printf("test%v\n", sendDone)
				break block
			}
		}
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := mainGenerater(done)

	for i := 0; i <= 10; i++ {
		fmt.Println(ch)
	}

	done <- struct{}{}
	//close(done)
	time.Sleep(1 * time.Second)
	fmt.Println(runtime.NumGoroutine())
	//for v := range ch {
	//	fmt.Printf("%v\n", v)
	//}
}
