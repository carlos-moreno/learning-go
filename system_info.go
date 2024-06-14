package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOOS)
	time.Sleep(time.Second * 5)
	fmt.Println(runtime.GOARCH)
}
