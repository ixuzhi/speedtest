#!/bin/bash
#https://github.com/didi/sgt/blob/master/control

buildBranch=$(git rev-parse --abbrev-ref HEAD)
commitHash=$(git rev-parse --short HEAD)
buildDate=$(date "+%Y-%m-%d-%H:%M:%S")
commitDate=$(git show -s --format=%cd --date=format:%Y:%m:%d_%H:%M:%S)

run_linux() {
	go run speedtest.go	
}

case $1 in
*)
	run_linux
	#echo "Usage:$0 {env | android | linux}"
	;;
esac
