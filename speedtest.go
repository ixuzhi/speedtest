package main

import (
	"fmt"
	//"github.com/ixuzhi/speedtest/speedtest"
	"github.com/ixuzhi/speedtest/speedtest"
)

func main() {
	var clientInfo speedtest.ClientInfo
	fmt.Println("speedtest init.")
	clientInfo, err := speedtest.GetIPAndLatLon()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", clientInfo)
	var Servers speedtest.ServerList
	Servers, err = speedtest.GetSpeedTestServersList()
	if err != nil {
		fmt.Println("speedtest.GetSpeedTestServersList:" + err.Error())
		return
	} else {
		fmt.Println("len:", len(Servers.ServersInfo))
		//for k,v:=range ServersInfo.ServersInfo{
		//	fmt.Printf("%v,%+v\n",k,v)
		//}
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
}
