package main

import "fmt"

func main() {
	s := "ascii éáíôêã"

	// impressao de valores por caracter
	for _, v := range s {
		fmt.Printf("%v - %b - %T - %#U - %#x\n", v, v, v, v, v)
	}

	fmt.Println()

	// impressao de valores por byte
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %b - %T - %#U - %#x\n", s[i], s[i], s[i], s[i], s[i])
	}
}
