package main

import (
	"net/http"
	"fmt"
)

func main() {
	resp, error := http.Get("https://api.caiyunapp.com/v2/TAkhjf8d1nlSlspN/121.6544,25.1552/realtime.json ")
	if error != nil {
		//deal with error
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
