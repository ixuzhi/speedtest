package speedtest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

type IpApiCoClientInfo struct {
	Ip                 string  `json:"ip"`
	City               string  `json:"city"`
	Region             string  `json:"region"`
	RegionCode         string  `json:"region_code"`
	Country            string  `json:"country"`
	CountryName        string  `json:"country_name"`
	ContinentCode      string  `json:"continent_code"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Timezone           string  `json:"timezone"`
	UtcOffset          string  `json:"utc_offset"`
	CountryCallingCode string  `json:"country_calling_code"`
	Currency           string  `json:"currency"`
	Languages          string  `json:"languages"`
	Asn                string  `json:"asn"`
	org                string  `json:"Org"`
}

/*
{
    "ip": "171.43.253.189",
    "city": "Wuhan",
    "region": "Hubei",
    "region_code": "HB",
    "country": "CN",
    "country_name": "China",
    "continent_code": "AS",
    "in_eu": false,
    "postal": null,
    "latitude": 30.5856,
    "longitude": 114.2665,
    "timezone": "Asia/Shanghai",
    "utc_offset": "+0800",
    "country_calling_code": "+86",
    "currency": "CNY",
    "languages": "zh-CN,yue,wuu,dta,ug,za",
    "asn": "AS4134",
    "org": "No.31,Jin-rong Street"
}
*/
func GetIpapiCoClientInfo() (ClientInfo, error) {
	var clientinfo ClientInfo
	ipUrl := "https://ipapi.co/json/"
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
		return clientinfo, errors.New(err.Error())
	}
	var ipdata IpApiCoClientInfo
	if err := json.Unmarshal(data, &ipdata); err != nil {
		return clientinfo, errors.New(err.Error())
	} else {
		clientinfo.ClientIP = ipdata.Ip
		clientinfo.ClientLat = ipdata.Latitude
		clientinfo.ClientLon = ipdata.Longitude
		clientinfo.ClientFrom = "https://ipapi.co/json"
		return clientinfo, err
	}
}
