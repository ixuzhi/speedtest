package speedtest

import (
	"errors"
	"fmt"
)

type ClientInfo struct {
	ClientIP   string
	ClientLat  float64
	ClientLon  float64
	ClientFrom string
}

func GetIPAndLatLon() (clientinfo ClientInfo, err error) {

	clientinfo, err = GetIpApiComClientInfo()
	if err == nil {
		return clientinfo, nil
	}
	fmt.Println(err)
	clientinfo, err = GetIpLaClientInfo()
	if err == nil {
		return clientinfo, nil
	}
	fmt.Println(err)
	clientinfo, err = GetSpeedTestConfigClientInfo()
	if err == nil {
		return clientinfo, nil
	}
	fmt.Println(err)
	clientinfo, err = GetIpInfoIoClientInfo()
	if err == nil {
		return clientinfo, nil
	}
	clientinfo, err = GetIpapiCoClientInfo()
	if err == nil {
		return clientinfo, nil
	}
	fmt.Println(err)
	return clientinfo, errors.New("not find ClientInfo")
}

/*
https://go101.org/article/channel-use-cases.html#first-response-wins
*/
func GetClientInfo() {
	clientinfo := make(chan ClientInfo, 5)
	go GetIPAndLatLon1(clientinfo)
	go GetIPAndLatLon2(clientinfo)
	go GetIPAndLatLon3(clientinfo)
	go GetIPAndLatLon4(clientinfo)
	go GetIPAndLatLon5(clientinfo)
	fmt.Println("GetClientInfo:", <-clientinfo)
}

func GetIPAndLatLon1(client chan<- ClientInfo) {
	clientinfo, err := GetIpApiComClientInfo()
	if err == nil {
		client <- clientinfo
	}
}

func GetIPAndLatLon2(client chan<- ClientInfo) {
	clientinfo, err := GetIpLaClientInfo()
	if err == nil {
		client <- clientinfo
	}
}

func GetIPAndLatLon3(client chan<- ClientInfo) {
	clientinfo, err := GetSpeedTestConfigClientInfo()
	if err == nil {
		client <- clientinfo
	}
}

func GetIPAndLatLon4(client chan<- ClientInfo) {
	clientinfo, err := GetIpInfoIoClientInfo()
	if err == nil {
		client <- clientinfo
	}
}

func GetIPAndLatLon5(client chan<- ClientInfo) {
	clientinfo, err := GetIpapiCoClientInfo()
	if err == nil {
		client <- clientinfo
	}
}
