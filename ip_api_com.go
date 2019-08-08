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
		    "as": "AS4134 No.31,Jin-rong Street",
		    "city": "Shizishan",
		    "country": "China",
		    "countryCode": "CN",
		    "isp": "Chinanet",
		    "lat": 30.5465,
		    "lon": 114.342,
		    "org": "Chinanet HB",
		    "query": "171.43.253.189",
		    "region": "HB",
		    "regionName": "Hubei",
		    "status": "success",
		    "timezone": "Asia/Shanghai",
		    "zip": ""
		}
	*/
	ipUrl := "http://ip-api.com/json"
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
