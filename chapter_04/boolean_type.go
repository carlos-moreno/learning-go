package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) // zero value
	x := true
	fmt.Println(x)

	result := 10 > 20
	fmt.Println(result)

	result = 100 > 50
	fmt.Println(result)
}
