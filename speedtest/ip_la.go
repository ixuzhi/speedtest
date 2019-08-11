package speedtest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
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
func GetIpLaClientInfo() (ClientInfo, error) {
	var clientinfo ClientInfo
	ipUrl := "https://api.ip.la/en?json"
	req, err := http.NewRequest("GET", ipUrl, nil)
	if err != nil {
		//fmt.Println(err)
		return clientinfo, errors.New(err.Error())
	}
	client := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		//fmt.Println(err)
		return clientinfo, errors.New(err.Error())
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println(err)
		return clientinfo, errors.New(err.Error())
	}
	//fmt.Println(string(data))
	var ipdata IpLaIp
	if err := json.Unmarshal(data, &ipdata); err != nil {
		return clientinfo, errors.New(err.Error())
	} else {
		//fmt.Printf("IpLaIp;%+v\n", ipdata)
		clientinfo.ClientIP = ipdata.Ip
		clientinfo.ClientLat, err = strconv.ParseFloat(ipdata.Iplalocation.Latitude, 64)
		clientinfo.ClientLon, err = strconv.ParseFloat(ipdata.Iplalocation.Longitude, 64)
		clientinfo.ClientFrom = "https://api.ip.la/en?json"
		return clientinfo, err
	}
}
