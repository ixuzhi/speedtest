package speedtest

import (
	"errors"
	"fmt"
	"io"
	"net"
	"time"
)

/*
https://github.com/sivel/go-speedtest/blob/master/speedtest.go#L409
how to calc Download speed.
*/

func SpeedTestTcpDownload(serverHost string, size uint64) (float64, error) {
	var timeout = time.Second * 60
	var dataSize = 25000000
	var timeStart time.Time
	timeStart = time.Now()
	conn, err := net.DialTimeout("tcp", serverHost, timeout)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return 0, errors.New(err.Error())
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
	//fmt.Printf("%v,%v,%6.4f\n", dataSize, timeLapse, speed)
	return speed, nil
}
