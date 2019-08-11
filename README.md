| Name | CI/Travis | Description |  
| ---------|---------|--------- |  
| circleci | [![Build Status](https://circleci.com/gh/ixuzhi/speedtest.svg?style=svg)](https://circleci.com/gh/ixuzhi/speedtest) | circleci.com |  
| Travis CI | [![Travis CI](https://api.travis-ci.org/ixuzhi/speedtest.svg?branch=master)](https://travis-ci.org/ixuzhi/speedtest) | travis-ci.org |
| godoc |[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/ixuzhi/speedtest/speedtest)| godoc.org |

## how to calculate network speed?

### en  
```
1. I hava to know my really ip and lontitude latitude.

2. I hava to know some speedtest servers for speedtest.

3. calc distance for my computer to the speedtest server,and find the best near server for test.

4. calc lantency for my computer to the speedtest server,and find the best low lantency server for test.

5. upload and download test for calc network speed.find the biggest for the result.
```

### zh
```
1、找到电脑所在的外网ip地址和经纬度

2、找到测速服务器列表

3、通过经纬度计算电脑所在经纬度到测速服务器的距离

4、选取距离最近的服务器，计算个人电脑到测速服务器的延迟

5、选取延迟最低的测速服务器，通过http或者tcp方式，上传测试、下载测试（通过上传、或下载指定大小字节数的文件，记下所需时间，即可计算出上传下载的速度，即所需的带宽）
```

```go
go run v0/ip_taobao_com.go
go run v0/ip_api_com.go
go run v0/ip_la.go
go run v0/speedtest_net_speedtest_config.go
go run v0/speedTestServer.go

go run v0/calcDistanceBylonAndLat.go
go run v0/calcLatency.go
go run v0/speedtestHttpDownload.go
go run v0/speedtestHttpUpload.go
```

[goland how to download and upload](https://progolang.com/how-to-download-files-in-go/)
[speedtest download](https://github.com/surol/speedtest-cli/blob/master/speedtest/download.go#L18)
[speedtest upload](https://github.com/surol/speedtest-cli/blob/master/speedtest/upload.go#L46)
