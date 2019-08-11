package speedtest

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func (servers *ServerList) calcLatency() {
	var timeout = time.Duration(time.Second * 3)
	server := servers.ServersInfo
	//var server []ServerInfo
	if len(servers.ServersInfo) >= 10 {
		server = servers.ServersInfo[:10]
	} else {
		server = servers.ServersInfo[:len(servers.ServersInfo)]
	}

	for k, v := range server {
		conn, err := net.DialTimeout("tcp", v.HostUrl, timeout)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		defer conn.Close()

		conn.Write([]byte("HI\n"))
		hello := make([]byte, 1024)
		conn.Read(hello)

		sum := time.Duration(0)
		for j := 0; j < 3; j++ {
			resp := make([]byte, 1024)
			start := time.Now()
			conn.Write([]byte(fmt.Sprintf("PING %d\n", start.UnixNano()/1000000)))
			conn.Read(resp)
			total := time.Since(start)
			sum += total
		}
		latency := sum / 3
		//fmt.Println("latency:", latency)
		servers.ServersInfo[k].Latency = latency.Seconds()
	}
}

func (servers *ServerList) SortByLatency() {
	sort.Slice(servers.ServersInfo[0:10], func(i, j int) bool {
		x := servers.ServersInfo[i].Latency
		y := servers.ServersInfo[j].Latency
		if x == 0 || y == 0 {
			return false
		} else {
			return servers.ServersInfo[i].Latency < servers.ServersInfo[j].Latency
		}
	})
}
