package main

import (
	"fmt"
	"net/http"
	"sync"
)

var temp sync.WaitGroup

var urls = []string{
	"http://www.golang.org/",
	"http://www.sweetbeecr.com",
	"https://www.baidu.com",
}

func main() {
	for _, url := range urls {
		temp.Add(1)
		go func(url string) {
			defer temp.Done()

			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(resp.Status, " ", url)
			} else {
				fmt.Printf("wrong!")
			}
		}(url)
	}
	temp.Wait()
}
