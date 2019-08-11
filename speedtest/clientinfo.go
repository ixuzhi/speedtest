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
