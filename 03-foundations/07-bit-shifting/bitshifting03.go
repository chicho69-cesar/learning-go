package main

import (
	"fmt"
)

func main() {
	s := 6
	fmt.Printf("%d\t\t%b\n", s, s)
	s = s << 1
	fmt.Printf("%d\t\t%b\n", s, s)
	s = s << 1
	fmt.Printf("%d\t\t%b\n", s, s)

	kb := 1024
	mb := 1024 * 1024
	gb := 1024 * 1024 * 1024
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
