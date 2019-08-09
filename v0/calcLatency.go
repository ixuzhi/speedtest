package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	var timeout =time.Duration(time.Second*30)
	var servers = []string{
		"speedtest2.hb.chinamobile.com:8080",
		"vipspeedtest1.wuhan.net.cn:8080",
		"vipspeedtest4.wuhan.net.cn:8080",
		"113.57.249.2:8080",
		"speedtest.tkfl.cc:8080",
	}
	if len(servers) >= 5 {
		servers = servers[:5]
	} else {
		servers = servers[:len(servers)]
	}

	for _, serverHost := range servers {
		//addr, err := net.ResolveTCPAddr("tcp", serverHost)
		//if err != nil {
		//	fmt.Printf("%s\n", err.Error())
		//	continue
		//}

		//conn, err := net.DialTCP("tcp", nil, addr)
		conn,err:=net.DialTimeout("tcp",serverHost,timeout)
		//conn, err := net.DialTimeout("tcp", server.speedtest.Source, addr, timeout)
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
		fmt.Println("latency:", latency)
	}

}
