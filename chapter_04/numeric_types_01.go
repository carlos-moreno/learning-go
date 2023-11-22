package main

import "fmt"

func main() {
	a := "e"
	b := "é"
	c := "\u9999"
	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)
	fmt.Printf("%v, %v, %v\n", d, e, f)
}
