package speedtest

import (
	"bytes"
	"errors"
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

func SpeedTestHttpUpload(url string, size uint64) (float64, error) {
	//var url = "https://vipspeedtest1.wuhan.net.cn:8080/upload"
	uploadUrl := fmt.Sprintf("http://%s/upload", url)
	//fmt.Println(uploadUrl)
	var uploadSize = size
	postData := make([]byte, uploadSize)
	dataSize, err := rand.Read(postData)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	timeStart := time.Now()
	req, err := http.NewRequest("POST", uploadUrl, bytes.NewReader(postData))
	if err != nil {
		return 0, errors.New(err.Error())
	}
	client := http.Client{
		Timeout: time.Second * 60,
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	timeLapse := time.Since(timeStart)
	timeCost := float64(timeLapse)
	speed := float64(dataSize) * 8 * float64(time.Second) / float64(timeCost) / 1000 / 1000
	return speed, nil
}
