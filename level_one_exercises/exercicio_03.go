package main

import "fmt"

var (
	x = 42
	y = "James Bond"
	z = true
)

func main() {
	s := fmt.Sprintf("%v\n%v\n%v", x, y, z)
	fmt.Println(s)
}
