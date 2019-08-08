package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
func main() {
	var FileName string = "server.dat"
	if IsExist(FileName) {
		fmt.Printf("%s exist\n", FileName)
	} else {
		fmt.Println("not exist")
		ipUrl := "https://www.speedtest.net/speedtest-servers.php"
		req, err := http.NewRequest("GET", ipUrl, nil)
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
			return
		} else {
			ioutil.WriteFile("server.dat", data, 0755)
			fmt.Println("ok")
		}
	}

}
