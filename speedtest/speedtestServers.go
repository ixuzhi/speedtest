package speedtest

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

/*
<?xml version="1.0" encoding="UTF-8"?>
<settings>
<servers>
<server url="http://speedtest.oppinord.no:8080/speedtest/upload.php" lat="70.6632" lon="23.6817" \
name="Hammerfest" country="Norway" cc="NO" sponsor="Hammerfest Energi Bredbånd AS" id="21239" \
host="speedtest.oppinord.no:8080" />
</servers>
</settings>
*/

//https://github.com/sivel/go-speedtest/blob/master/speedtest.go#L301
type ServerList struct {
	ServersInfo []ServerInfo `xml:"servers>server"`
}

type ServerInfo struct {
	Url      string  `xml:"url,attr"`
	Lat      float64 `xml:"lat,attr"`
	Lon      float64 `xml:"lon,attr"`
	Name     string  `xml:"name,attr"`
	Country  string  `xml:"country,attr"`
	Cc       string  `xml:"cc,attr"`
	Sponsor  string  `xml:"sponsor,attr"`
	Id       string  `xml:"id,attr"`
	HostUrl  string  `xml:"host,attr"`
	Distance float64
	Latency  float64
}

func IsExist(f string) bool {
	fmt.Println(f)
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
func GetSpeedTestServersList() (ServerList, error) {
	var ServersInfo ServerList
	var FileName = "./speedTestServers.dat"
	var data []byte
	var err error
	if IsExist(FileName) {
		fmt.Printf("%s exist\n", FileName)
		data, err = ioutil.ReadFile(FileName)
		if err != nil {
			//fmt.Println(err)
			return ServersInfo, errors.New("ioutil.ReadFile:" + err.Error())
		}
	} else {
		fmt.Println("not exist")
		//ipUrl := "https://www.speedtest.net/speedtest-servers.php"
		ipUrl := "https://c.speedtest.net/speedtest-servers-static.php"
		req, err := http.NewRequest("GET", ipUrl, nil)
		if err != nil {
			//fmt.Println(err)
			return ServersInfo, errors.New("http.NewRequest:" + err.Error())
		}
		client := http.Client{
			Timeout: time.Second * 90,
		}
		resp, err := client.Do(req)
		if err != nil {
			//fmt.Println(err)
			return ServersInfo, errors.New("client.Do:" + err.Error())
		}
		defer resp.Body.Close()

		data, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			//fmt.Println(err)
			return ServersInfo, errors.New(err.Error())
		} else {
			ioutil.WriteFile(FileName, data, 0755)
			fmt.Println("ok")
		}
	}

	//ServersInfo.ServersInfo
	if err := xml.Unmarshal(data, &ServersInfo); err != nil {
		//fmt.Println(err)
		return ServersInfo, errors.New(err.Error())
	} else {
		//fmt.Println(len(ServersInfo.ServersInfo))
		//for k, v := range ServersInfo.ServersInfo {
		//	fmt.Printf("%d,%+v\n", k, v)
		//}
		return ServersInfo, nil
	}
}

func (servers *ServerList) GetClosestSpeedTestServers(clientinfo ClientInfo) {
	for k, v := range servers.ServersInfo {
		latLonTestServer := pos{
			φ: v.Lat, // latitude, radians
			ψ: v.Lon, // longitude, radians
		}
		distance := hsDist(degPos(clientinfo.ClientLat, clientinfo.ClientLon), degPos(latLonTestServer.φ, latLonTestServer.ψ))
		//fmt.Println(k,latLon,latLonTestServer,distance,v.HostUrl)
		servers.ServersInfo[k].Distance = distance
	}
	servers.SortByDistance()
	//for k, v := range servers.ServersInfo[0:10] {
	//	fmt.Printf("11:%v,%+v,%v\n", k, v.Distance, v.HostUrl)
	//}
	servers.calcLatency()
	servers.SortByLatency()
	//for k, v := range servers.ServersInfo[0:10] {
	//	fmt.Printf("22:%v,%+v,%v\n", k, v.Distance, v.HostUrl)
	//}
}
