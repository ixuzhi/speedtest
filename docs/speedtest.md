https://github.com/sivel/go-speedtest/blob/master/speedtest.go#L363  

## how to get ip and lan lon by http request?

   [best-commercial-ip-geolocation-api](https://medium.com/@ipdata_co/what-is-the-best-commercial-ip-geolocation-api-d8195cda7027#6605)  
   [api.ip.la](https://api.ip.la/en?json)  
   [ip-api.com](http://ip-api.com/json)  
   [speedtest-config.php](https://www.speedtest.net/speedtest-config.php)  
   [ip.taobao.com](http://ip.taobao.com/service/getIpInfo2.php?ip=myip)  
    
## how to select nearest speedtest server by lat and lon.
https://rosettacode.org/wiki/Haversine_formula  
https://rosettacode.org/wiki/Haversine_formula#Go  

## how to select speedtest server，calc latency
[Reverse Engineering the Speedtest.net Protocol](https://gist.github.com/sdstrowes/411fca9d900a846a704f68547941eb97)

[reverse_engineering_the_speedtest_net_protocol](https://web.archive.org/web/20141216073338/https://gkbrk.com/blog/read?name=reverse_engineering_the_speedtest_net_protocol)

[How does the test itself work? How is the result calculated?](https://support.speedtest.net/hc/en-us/articles/203845400-How-does-the-test-itself-work-How-is-the-result-calculated-)
```
1. select speedtest server by https://www.speedtest.net/speedtest-servers.php
   get the host field,get the tcpaddr
2、establish tcp connect
    
    send: "HI\n"
    recv: "HELLO 2.6 (2.6.9) 2019-02-20.2246.62a8e21"
    send: "PING 1559029313650"
    recv: "PONG 1559028780098"
    
    send and recv multi time and calc the latency from your computer to speedtest server. 
```

## how to sort by distance
[distance sort](https://github.com/sivel/go-speedtest/blob/master/speedtest.go#L306)
