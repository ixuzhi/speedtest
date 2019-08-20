package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type SpeedTestDotNetServerInfo struct {
	Url             string `json:"url"`
	Lat             string `json:"lat"`
	Lon             string `json:"lon"`
	Distance        string `json:"distance"`
	Name            string `json:"name"`
	Country         string `json:"country"`
	Cc              string `json:"cc"`
	Sponsor         string `json:"sponsor"`
	Id              string `json:"id"`
	Preferred       int32  `json:"id"`
	HttpsFunctional int32  `json:"https_functional"`
	Host            string `json:"host"`
}

func GetSpeedTestServersList_speedtest_dotnet() []SpeedTestDotNetServerInfo {
	ipUrl := "https://www.speedtest.net/api/js/servers?engine=js&https_functional=1&limit=5"
	req, err := http.NewRequest("GET", ipUrl, nil)
	if err != nil {
		return nil
	}
	client := http.Client{
		Timeout: time.Second * 90,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	} else {
		fmt.Println("ok")
	}
	//fmt.Println(string(data))
	var servers []SpeedTestDotNetServerInfo
	if err := json.Unmarshal(data, &servers); err != nil {

	}
	fmt.Println(len(servers))
	fmt.Println(servers)
	return servers
}

//
//func main() {
//	GetSpeedTestServersList_speedtest_dotnet()
//}
