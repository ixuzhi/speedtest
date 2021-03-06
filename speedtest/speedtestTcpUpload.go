package speedtest

import (
	"errors"
	"fmt"
	"net"
	"time"
)

/*
https://github.com/sivel/go-speedtest/blob/master/speedtest.go#L502
how to calc tcp Upload speed.
*/

func SpeedTestTcpUpload(serverHost string, size uint64) (float64, error) {
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

	var give int
	var dataSend []byte
	remaining := dataSize

	for remaining > 0 && time.Since(timeStart).Seconds() < float64(timeout) {
		if remaining > 100000 {
			give = 100000
		} else {
			give = remaining
		}
		header := []byte(fmt.Sprintf("UPLOAD %d 0\n", give))
		if give-len(header) > 0 {
			dataSend = make([]byte, give-len(header))
			conn.Write(header)
			conn.Write(dataSend)
		} else {

		}
		up := make([]byte, 24)
		conn.Read(up)
		remaining -= give
	}
	timeLapse := time.Since(timeStart)
	timeCost := float64(timeLapse)
	speed := float64(dataSize) * 8 * float64(time.Second) / float64(timeCost) / 1000 / 1000
	//fmt.Println(speed, timeLapse)
	return speed, nil
}
