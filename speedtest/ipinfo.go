package speedtest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type IpInfoIoClientInfo struct {
	Ip     string `json:"ip"`
	City   string `json:"city"`
	Region string `json:"region"`
	Loc    string `json:"loc"`
	Org    string `json:"org"`
}

/*
{
  "ip": "171.43.253.189",
  "city": "Jinkou",
  "region": "Hubei",
  "country": "CN",
  "loc": "30.3401,114.1299",
  "org": "AS4134 CHINANET-BACKBONE"
}
*/
func GetIpInfoIoClientInfo() (ClientInfo, error) {
	var clientinfo ClientInfo
	ipUrl := "http://ipinfo.io/?token=2c1be40be8a245"
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
	var ipdata IpInfoIoClientInfo
	if err := json.Unmarshal(data, &ipdata); err != nil {
		return clientinfo, errors.New(err.Error())
	} else {
		//fmt.Printf("IpLaIp;%+v\n", ipdata)
		clientinfo.ClientIP = ipdata.Ip
		splitLonlat := strings.Split(ipdata.Loc, ",")
		if len(splitLonlat) == 2 {
			clientinfo.ClientLat, err = strconv.ParseFloat(splitLonlat[0], 64)
			clientinfo.ClientLon, err = strconv.ParseFloat(splitLonlat[1], 64)
		}
		return clientinfo, err
	}
}
