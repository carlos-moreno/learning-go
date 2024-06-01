package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOOS)
	time.Sleep(time.Second * 11)
	fmt.Println(runtime.GOARCH)
}
