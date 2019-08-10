package main

import (
	"fmt"
	//"github.com/ixuzhi/speedtest/speedtest"
	"./speedtest"
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
}
