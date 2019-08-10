package speedtest

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

/*
https://www.speedtest.net/speedtest-config.php
<settings>
<client ip="171.43.253.189" lat="30.5856" lon="114.2665" isp="China Telecom Hubei" isprating="3.7" rating="0" ispdlavg="0" ispulavg="0" loggedin="0" country="CN"/>
</settings>
*/
type Client struct {
	Ip        string `xml:"ip,attr"`
	Lat       string `xml:"lat,attr"`
	Lon       string `xml:"lon,attr"`
	Isp       string `xml:"isp,attr"`
	Isprating string `xml:"isprating,attr"`
	Country   string `xml:"country,attr"`
}

type SpeedTestClientConfig struct {
	Client Client `xml:"client"`
}

func GetSpeedTestConfigClientInfo() (ClientInfo, error) {
	var clientinfo ClientInfo
	ipUrl := "https://www.speedtest.net/speedtest-config.php"
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
	var ipdata SpeedTestClientConfig
	if err := xml.Unmarshal(data, &ipdata); err != nil {
		return clientinfo, errors.New(err.Error())
	} else {
		//fmt.Printf("SpeedTestClientConfig;%+v\n", ipdata)
		clientinfo.ClientIP = ipdata.Client.Ip
		clientinfo.ClientLat, err = strconv.ParseFloat(ipdata.Client.Lat, 64)
		clientinfo.ClientLon, err = strconv.ParseFloat(ipdata.Client.Lon, 64)
		return clientinfo, nil
	}
}
