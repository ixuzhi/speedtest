package speedtest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type TaobaoData struct {
	Ip         string `json:"ip"`
	Country    string `json:"country"`
	Area       string `json:"area"`
	Region     string `json:"region"`
	City       string `json:"city"`
	County     string `json:"county"`
	Isp        string `json:"isp"`
	Country_id string `json:"country_id"`
	Area_id    string `json:"area_id"`
	Region_id  string `json:"region_id"`
	City_id    string `json:"city_id"`
	County_id  string `json:"county_id"`
	Isp_id     string `json:"isp_id"`
}

type TaobaoIp struct {
	Data TaobaoData `json:"data"`
	code int32      `json:"code"`
}

/*
	{
	    "code": 0,
	    "data": {
	        "ip": "171.43.253.189",
	        "country": "中国",
	        "area": "",
	        "region": "湖北",
	        "city": "武汉",
	        "county": "XX",
	        "isp": "电信",
	        "country_id": "CN",
	        "area_id": "",
	        "region_id": "420000",
	        "city_id": "420100",
	        "county_id": "xx",
	        "isp_id": "100017"
	    }
	}
*/
func GetIpTaoBaoClientInfo() (ClientInfo, error) {
	var clientinfo ClientInfo
	ipUrl := "http://ip.taobao.com/service/getIpInfo2.php?ip=myip"
	req, err := http.NewRequest("POST", ipUrl, nil)
	if err != nil {
		fmt.Println(err)
		return clientinfo, errors.New(err.Error())
	}
	client := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return clientinfo, errors.New(err.Error())
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return clientinfo, errors.New(err.Error())
	}
	fmt.Println(string(data))
	var ipdata TaobaoIp
	if err := json.Unmarshal(data, &ipdata); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("TaobaoIp;%+v\n", ipdata)
		clientinfo.ClientIP = ipdata.Data.Ip
		clientinfo.ClientFrom = "http://ip.taobao.com/service/getIpInfo2.php?ip=myip"
		//clientinfo.ClientLat = ipdata.Data.
	}
	return clientinfo, errors.New(err.Error())
}
