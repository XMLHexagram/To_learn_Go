package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type weather struct {
	status      string
	lang        string
	api_version string
}

func main() {
	fmt.Println(b)
	var temp weather

	res, err := http.Get("https://api.caiyunapp.com/v2/TAkhjf8d1nlSlspN/121.6544,25.1552/realtime.json")
	checkError(err)
	data, err := ioutil.ReadAll((res.Body))
	checkError(err)
	json.Unmarshal(data, &temp)
	fmt.Println(temp)
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}
