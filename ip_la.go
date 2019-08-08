package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	/*
		{
		    "ip": "171.43.253.189",
		    "location": {
		        "city": "Wuhan",
		        "country_code": "CN",
		        "country_name": "China",
		        "latitude": "30.572399",
		        "longitude": "114.279121",
		        "province": "Hubei"
		    }
		}
	*/
	ipUrl := "https://api.ip.la/en?json"
	req, err := http.NewRequest("GET", ipUrl, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
