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
		    "code": 0,
		    "data": {
		        "ip": "171.43.253.189",
		        "country": "中国",
		        "area": "",
		        "region": "湖北",
		        "city": "武汉",
		        "county": "XX",
		        "isp": "电信",
		        "country_id": "CN",
		        "area_id": "",
		        "region_id": "420000",
		        "city_id": "420100",
		        "county_id": "xx",
		        "isp_id": "100017"
		    }
		}
	*/
	ipUrl := "http://ip.taobao.com/service/getIpInfo2.php?ip=myip"
	req, err := http.NewRequest("POST", ipUrl, nil)
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
