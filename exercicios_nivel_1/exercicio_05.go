package main

import "fmt"

type xpto int

var x xpto
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%T\n", y)
}
