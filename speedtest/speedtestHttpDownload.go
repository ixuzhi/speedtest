package speedtest

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*
https://vipspeedtest1.wuhan.net.cn:8080/download?nocache=c20185fc-18b8-4f65-813b-9529f7417961&size=25000000&guid=fde67500-60af-4d78-b13f-5c53d4988817
https://vipspeedtest1.wuhan.net.cn:8080/upload?nocache=fcc54550-fac3-42e2-b753-e34e2c442a6f&guid=fde67500-60af-4d78-b13f-5c53d4988817

how to calc download speed.
https://github.com/sivel/go-speedtest/blob/master/speedtest.go#L679
https://github.com/surol/speedtest-cli/blob/master/speedtest/download.go#L92

*/
func SpeedTestHttpDownload(url string, size uint64) (float64, error) {
	//var url = "https://vipspeedtest1.wuhan.net.cn:8080/download?size=25000000"
	downloadUrl := fmt.Sprintf("http://%s/download?size=%d", url, size)
	timeStart := time.Now()
	req, err := http.NewRequest("GET", downloadUrl, nil)
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

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.New(err.Error())
	}
	timeLapse := time.Since(timeStart)
	byteSizeDownload := len(data)
	timeCost := float64(timeLapse)

	speed := float64(byteSizeDownload) * 8 * float64(time.Second) / float64(timeCost) / 1000 / 1000
	//fmt.Println(len(data), timeLapse, speed, t)
	return speed, nil
}
