package speedtest

type ClientInfo struct {
	ClientIP  string
	ClientLat float64
	ClientLon float64
}

func GetIPAndLatLon() (clientinfo ClientInfo, err error) {
	clientinfo, err = GetIpApiComClientInfo()
	if err == nil {
		return clientinfo, nil
	}
	clientinfo, err = GetIpLaClientInfo()
	if err == nil {
		return clientinfo, nil
	}
	clientinfo, err = GetSpeedTestConfigClientInfo()
	if err == nil {
		return clientinfo, nil
	}
	return
}
