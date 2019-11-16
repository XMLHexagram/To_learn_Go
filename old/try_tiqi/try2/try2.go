package main

import (
	"encoding/json"
)

type server struct {
	ServerName string
	ServerIP   string
}

type serverslice struct {
	Servers []server
}

func main() {
	var s serverslice

	// 模拟传输的Json数据
	str := `{
                "servers": [
                    {
                        "serverName": "Shanghai_VPN",
                        "serverIP": "127.0.0.1"
                    }, {
                        "serverName": "Beijing_VPN",
                        "serverIP": "127.0.0.2"
                    }
                ]
            }`

	// 解析字符串为Json
	json.Unmarshal([]byte(str), &s)

	// 遍历Json
	for key, val := range s.Servers {

		// 打印索引和其他数据
		println("Key：", key, "\tName：", val.ServerName, "\tIP：", val.ServerIP)
	}
}
