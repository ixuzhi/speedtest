package main

import (
	"fmt"
	//"github.com/ixuzhi/speedtest/speedtest"
	"github.com/ixuzhi/speedtest/speedtest"
)

var (
	dataSize      uint64 = 2500000
	Servers       speedtest.ServerList
	testServerLen int64 = 1
	clientInfo    speedtest.ClientInfo
)

func getspeed() {
	if len(Servers.ServersInfo) > 0 {
		Servers.GetClosestSpeedTestServers(clientInfo)
		for k, v := range Servers.ServersInfo[0:10] {
			fmt.Printf("|111:%-4d|%-10.4f|%-10.4f|%-30s\n", k, v.Latency*1000, v.Distance, v.HostUrl)
		}
	} else {
		fmt.Println("len Servers.ServersInfo ==0")
		return
	}

	for k, v := range Servers.ServersInfo[0:10] {
		fmt.Printf("|%-4d|%-10.4f|%-10.4f|%-30s\n", k, v.Latency*1000, v.Distance, v.HostUrl)
	}
	for k, v := range Servers.ServersInfo[0:testServerLen] {
		up, err := speedtest.SpeedTestTcpUpload(v.HostUrl, dataSize)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("|%-4d |Upload tcp:%-6.2f\n%+v\n", k, up, v)
	}

	for k, v := range Servers.ServersInfo[0:testServerLen] {
		up, err := speedtest.SpeedTestTcpDownload(v.HostUrl, dataSize)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("|%-4d |Download tcp:%-6.2f\n%+v\n", k, up, v)
	}

	for k, v := range Servers.ServersInfo[0:testServerLen] {
		up, err := speedtest.SpeedTestHttpDownload(v.HostUrl, dataSize)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("|%-4d |Download http:%-6.2f\n%+v\n", k, up, v)
	}
	for k, v := range Servers.ServersInfo[0:testServerLen] {
		up, err := speedtest.SpeedTestHttpUpload(v.HostUrl, dataSize)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("|%-4d |Upload http:%-6.2f\n%+v\n", k, up, v)
	}
}

func getspeed2() {

	speedtestDotNetServers := speedtest.GetSpeedTestServersList_speedtest_dotnet()
	if speedtestDotNetServers == nil {
		return
	} else {
		for k, v := range speedtestDotNetServers[0:len(speedtestDotNetServers)] {
			up, err := speedtest.SpeedTestTcpUpload(v.Host, dataSize)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("|%-4d |Upload tcp:%-6.2f\n%+v\n", k, up, v)

			tcpDownload, err := speedtest.SpeedTestTcpDownload(v.Host, dataSize)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("|%-4d |Download tcp:%-6.2f\n%+v\n", k, tcpDownload, v)

			httpDownload, err := speedtest.SpeedTestHttpDownload(v.Host, dataSize)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("|%-4d |Download http:%-6.2f\n%+v\n", k, httpDownload, v)

			httpUpload, err := speedtest.SpeedTestHttpUpload(v.Host, dataSize)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("|%-4d |Upload http:%-6.2f\n%+v\n", k, httpUpload, v)
		}
	}
}

func main() {
	var err error
	fmt.Println("speedtest init.")
	speedtest.EnvInfo()
	clientInfo, err = speedtest.GetIPAndLatLon()
	if err != nil {
		fmt.Println(err)
		return
	}
	speedtest.GetClientInfo()
	fmt.Printf("%+v\n", clientInfo)

	Servers, err = speedtest.GetSpeedTestServersList()
	if err != nil {
		fmt.Println("speedtest.GetSpeedTestServersList:" + err.Error())
		return
	} else {
		fmt.Println("len:", len(Servers.ServersInfo))
	}
	getspeed()
	getspeed2()
}
