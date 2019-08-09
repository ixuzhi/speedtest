https://github.com/sivel/go-speedtest/blob/master/speedtest.go#L363  

## how to select speedtest server，calc latency
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
