package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

/*
https://vipspeedtest1.wuhan.net.cn:8080/download?nocache=c20185fc-18b8-4f65-813b-9529f7417961&size=25000000&guid=fde67500-60af-4d78-b13f-5c53d4988817
https://vipspeedtest1.wuhan.net.cn:8080/upload?nocache=fcc54550-fac3-42e2-b753-e34e2c442a6f&guid=fde67500-60af-4d78-b13f-5c53d4988817

how to calc upload speed.


*/
func speedtestHttpUpload() {
	var url = "https://vipspeedtest1.wuhan.net.cn:8080/upload"
	var uploadSize = 25000000
	postData := make([]byte, uploadSize)
	datasize, err := rand.Read(postData)
	if err != nil {
		fmt.Println(err, datasize)
		return
	}

	timeStart := time.Now()
	req, err := http.NewRequest("POST", url, bytes.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return
	}
	client := http.Client{
		Timeout: time.Second * 60,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("recv:", string(data))
	timeEnd := time.Now()
	t := timeEnd.Sub(timeStart)

	timeLapse := time.Since(timeStart)
	timeCost := float64(timeLapse)

	speed := float64(datasize) * 8 * float64(time.Second) / float64(timeCost) / 1000 / 1000
	fmt.Println(len(data), timeLapse, speed, t)
}

func main() {
	speedtestHttpUpload()
}
