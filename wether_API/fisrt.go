package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./server"
)

func main() {
	go server.MainOut()
	go jieXiURL()
}

func jieXiURL() {
	var temp weather
	url := "https://api.caiyunapp.com/v2/TAkhjf8d1nlSlspN/121.6544,25.1552/realtime.json"
	json.Unmarshal(getJS(url), &temp)
	fmt.Println(temp)
	fmt.Printf("\n\n")
	temp.printWeather()
}

func (w weather) printWeather() {
	fmt.Println(w.Status)
	fmt.Println(w.Lang)
	fmt.Println(w.Unit)
}

func getJS(URL string) []byte {
	res, err := http.Get(URL)
	checkError(err)
	data, err := ioutil.ReadAll((res.Body))
	checkError(err)
	return data
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}

type weather struct {
	Status     string    `json:"status"`
	Lang       string    `json:"lang"`
	Unit       string    `json:"unit"`
	ServerTime int       `json:"server_time"`
	Location   []float64 `json:"location"`
	APIStatus  string    `json:"api_status"`
	Tzshift    int       `json:"tzshift"`
	APIVersion string    `json:"api_version"`
	Result     struct {
		Status        string  `json:"status"`
		O3            float64 `json:"o3"`
		Co            float64 `json:"co"`
		Temperature   float64 `json:"temperature"`
		Pm10          int     `json:"pm10"`
		Skycon        string  `json:"skycon"`
		Cloudrate     float64 `json:"cloudrate"`
		Precipitation struct {
			Nearest struct {
				Status    string  `json:"status"`
				Distance  float64 `json:"distance"`
				Intensity float64 `json:"intensity"`
			} `json:"nearest"`
			Local struct {
				Status     string `json:"status"`
				Intensity  int    `json:"intensity"`
				Datasource string `json:"datasource"`
			} `json:"local"`
		} `json:"precipitation"`
		Dswrf       float64 `json:"dswrf"`
		Visibility  int     `json:"visibility"`
		Humidity    float64 `json:"humidity"`
		So2         float64 `json:"so2"`
		Ultraviolet struct {
			Index int    `json:"index"`
			Desc  string `json:"desc"`
		} `json:"ultraviolet"`
		Pres                float64 `json:"pres"`
		Aqi                 int     `json:"aqi"`
		Pm25                int     `json:"pm25"`
		No2                 float64 `json:"no2"`
		ApparentTemperature float64 `json:"apparent_temperature"`
		Comfort             struct {
			Index int    `json:"index"`
			Desc  string `json:"desc"`
		} `json:"comfort"`
		Wind struct {
			Direction int     `json:"direction"`
			Speed     float64 `json:"speed"`
		} `json:"wind"`
	} `json:"result"`
}
