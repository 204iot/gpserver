package main

import (
	"log"

	"github.com/204iot/gpserver/server/cache"
	"github.com/go-resty/resty/v2"
)

func main() {
	client := resty.New()

	// 上传设备gps 信息
	// 43.6440258500,87.8466796900
	res, err := client.R().SetBody(cache.IOTDevice{
		ID:        "12",
		Longitude: "87.8466796900",
		Latitude:  "43.6440258500",
	}).Post("http://127.0.0.1:8080/api/v1/device")
	if err != nil {
		panic(err)
	}
	if res.StatusCode() != 200 {
		panic("Post fail")
	}

	// 获取所有设备信息
	if res, err = client.R().Get("http://127.0.0.1:8080/api/v1/devices"); err != nil {
		panic(err)
	}
	if res.StatusCode() != 200 {
		panic("Get fail")
	}
	log.Println(string(res.Body()))
}
