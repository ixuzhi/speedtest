package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type IpApiIp struct {
	As          string  `json:"as"`
	City        string  `json:"City"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Isp         string  `json:"isp"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Org         string  `json:"org"`
	Query       string  `json:"query"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	Status      string  `json:"status"`
	Timezone    string  `json:"timezone"`
	Zip         string  `json:"zip"`
}

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
	var ipdata IpApiIp
	if err := json.Unmarshal(data, &ipdata); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("ipapiip;%+v\n", ipdata)
	}
}
