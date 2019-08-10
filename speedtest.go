package main

import (
	"fmt"
	//"github.com/ixuzhi/speedtest/speedtest"
	speedtest "github.com/ixuzhi/speedtest/speedtest"
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
	var ServersInfo speedtest.ServerList
	ServersInfo, err = speedtest.GetSpeedTestServersList()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("len:", len(ServersInfo.Servers))
		//for k,v:=range ServersInfo.Servers{
		//	fmt.Printf("%v,%+v\n",k,v)
		//}
	}

}
