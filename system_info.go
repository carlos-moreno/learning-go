package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOOS)
	time.Sleep(time.Second * 6)
	fmt.Println(runtime.GOARCH)
}
