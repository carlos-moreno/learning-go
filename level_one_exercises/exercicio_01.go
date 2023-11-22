package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("x =", x, "y =", y, "z =", z)
	fmt.Printf("x = %+v\n", x)
	fmt.Printf("y = %+v\n", y)
	fmt.Printf("z = %+v\n", z)
}
