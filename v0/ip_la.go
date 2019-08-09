package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type IpLaLocation struct {
	City         string `json:"city"`
	Country_code string `json:"country_code"`
	Country_name string `json:"country_name"`
	Latitude     string `json:"latitude"`
	Longitude    string `json:"longitude"`
	Province     string `json:"province"`
}

type IpLaIp struct {
	Iplalocation IpLaLocation `json:"location"`
	Ip           string       `json:"ip"`
}

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
	var ipdata IpLaIp
	if err := json.Unmarshal(data, &ipdata); err != nil {

	} else {
		fmt.Printf("IpLaIp;%+v\n", ipdata)
	}
}
