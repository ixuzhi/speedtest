package main

import "fmt"

type ClientInfo struct {
	ClientIP  string
	ClientLat float64
	ClientLon float64
}

func getIPAndLatLon() (client ClientInfo, err error) {
	var clientInfo ClientInfo
	return clientInfo, nil
}

func main() {
	var clientInfo ClientInfo
	fmt.Println("speedtest init.")
	clientInfo, err := getIPAndLatLon()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", clientInfo)
}
