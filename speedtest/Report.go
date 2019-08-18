package speedtest

import (
	"fmt"
	"runtime"
)

func EnvInfo() {
	fmt.Println("env")
	fmt.Println("arch:	", runtime.GOARCH)
	fmt.Println("os:	", runtime.GOOS)
	fmt.Println("goroot: ", runtime.GOROOT())
	fmt.Println("numcpu: ", runtime.NumCPU())
}
