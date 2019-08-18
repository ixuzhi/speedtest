package main

import (
	"fmt"
	//"github.com/ixuzhi/speedtest/speedtest"
	"github.com/ixuzhi/speedtest/speedtest"
)

func main() {
	var clientInfo speedtest.ClientInfo
	var dataSize uint64 = 2500000
	fmt.Println("speedtest init.")
	speedtest.EnvInfo()
	clientInfo, err := speedtest.GetIPAndLatLon()
	if err != nil {
		fmt.Println(err)
		return
	}
	speedtest.GetClientInfo()
	fmt.Printf("%+v\n", clientInfo)

	var Servers speedtest.ServerList
	Servers, err = speedtest.GetSpeedTestServersList()
	if err != nil {
		fmt.Println("speedtest.GetSpeedTestServersList:" + err.Error())
		return
	} else {
		fmt.Println("len:", len(Servers.ServersInfo))
	}

	if len(Servers.ServersInfo) > 0 {
		Servers.GetClosestSpeedTestServers(clientInfo)
	} else {
		fmt.Println("len Servers.ServersInfo ==0")
		return
	}

	for k, v := range Servers.ServersInfo[0:10] {
		fmt.Printf("|%-4d|%-10.4f|%-10.4f|%-30s\n", k, v.Latency*1000, v.Distance, v.HostUrl)
	}
	for k, v := range Servers.ServersInfo[0:10] {
		up, err := speedtest.SpeedTestTcpUpload(v.HostUrl, dataSize)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("|%-4d |Upload tcp:%-6.2f\n%+v\n", k, up, v)
	}

	for k, v := range Servers.ServersInfo[0:10] {
		up, err := speedtest.SpeedTestTcpDownload(v.HostUrl, dataSize)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("|%-4d |Download tcp:%-6.2f\n%+v\n", k, up, v)
	}

	for k, v := range Servers.ServersInfo[0:10] {
		up, err := speedtest.SpeedTestHttpDownload(v.HostUrl, dataSize)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("|%-4d |Download http:%-6.2f\n%+v\n", k, up, v)
	}
	for k, v := range Servers.ServersInfo[0:10] {
		up, err := speedtest.SpeedTestHttpUpload(v.HostUrl, dataSize)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("|%-4d |Upload http:%-6.2f\n%+v\n", k, up, v)
	}
}
