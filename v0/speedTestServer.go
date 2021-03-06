package main

import (
	"encoding/xml"
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
	Servers []ServerInfo `xml:"servers>server"`
}

type ServerInfo struct {
	Url     string `xml:"url,attr"`
	Lat     string `xml:"lat,attr"`
	Lon     string `xml:"lon,attr"`
	Name    string `xml:"name,attr"`
	Country string `xml:"country,attr"`
	Cc      string `xml:"cc,attr"`
	Sponsor string `xml:"sponsor,attr"`
	Id      string `xml:"id,attr"`
	HostUrl string `xml:"host,attr"`
}

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
func main() {
	var FileName = "servers.dat"
	var data []byte
	var err error
	if IsExist(FileName) {
		fmt.Printf("%s exist\n", FileName)
		data, err = ioutil.ReadFile(FileName)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("not exist")
		ipUrl := "https://www.speedtest.net/speedtest-servers.php"
		req, err := http.NewRequest("GET", ipUrl, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		client := http.Client{
			Timeout: time.Second * 60,
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
		} else {
			ioutil.WriteFile(FileName, data, 0755)
			fmt.Println("ok")
		}
	}
	var ServersInfo ServerList
	//ServersInfo.Servers
	if err := xml.Unmarshal(data, &ServersInfo); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(len(ServersInfo.Servers))
		for k, v := range ServersInfo.Servers {
			fmt.Printf("%d,%+v\n", k, v)
		}
	}

}
