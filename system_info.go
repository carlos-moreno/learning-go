package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOOS)
	time.Sleep(time.Second * 9)
	fmt.Println(runtime.GOARCH)
}
