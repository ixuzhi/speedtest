package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

/*
https://github.com/sivel/go-speedtest/blob/master/speedtest.go#L409

how to calc Download speed.

*/
func speedtestTcpDownload() {
	var timeout = time.Second * 60
	var dataSize = 25000000
	var timeStart time.Time
	var servers = []string{
		"speedtest2.hb.chinamobile.com:8080",
		"vipspeedtest1.wuhan.net.cn:8080",
		"vipspeedtest4.wuhan.net.cn:8080",
		"113.57.249.2:8080",
	}
	for _, serverHost := range servers {
		timeStart = time.Now()
		conn, err := net.DialTimeout("tcp", serverHost, timeout)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		defer conn.Close()
		conn.Write([]byte("HI\n"))
		hello := make([]byte, 1024)
		conn.Read(hello)
		//fmt.Println(string(hello), len(hello))
		tmp := make([]byte, 1024)
		var out []int
		var readCount int
		remaining := dataSize

		for remaining > 0 && time.Since(timeStart).Seconds() < float64(timeout) {
			if remaining > 1000000 {
				readCount = 1000000
			} else {
				readCount = remaining
			}
			down := 0
			conn.Write([]byte(fmt.Sprintf("DOWNLOAD %d\n", readCount)))
			for down < readCount {
				n, err := conn.Read(tmp)
				if err != nil {
					if err != io.EOF {
						fmt.Printf("ERR: %v\n", err)
					}
					break
				}
				down += n
			}
			out = append(out, down)
			remaining -= down
		}

		timeLapse := time.Since(timeStart)
		speed := float64(dataSize) * 8 * float64(time.Second) / float64(timeLapse) / 1000 / 1000
		fmt.Printf("%v,%v,%6.4f\n", dataSize, timeLapse, speed)
	}
}

func main() {
	speedtestTcpDownload()
}
