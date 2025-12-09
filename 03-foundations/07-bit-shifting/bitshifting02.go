package main

import (
	"fmt"
)

const (
	_  = iota // 0
	kb = 1 << (iota * 10) // Recorremos 1(10) diez ceros a la izquierda
	mb = 1 << (iota * 10) // 10000000000 diez ceros a la izquierda
	gb = 1 << (iota * 10) // 100000000000000000000 diez ceros a la izquierda
)

func main() {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
