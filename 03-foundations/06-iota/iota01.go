package main

import (
	"fmt"
)

const (
	c0 = iota // Inicia en 0
	c1 = iota // 1 - Porque no ha aparecido la keyword const
	c2 = iota // 2
)

func main() {
	fmt.Println(c0)
	fmt.Println(c1)
	fmt.Println(c2)
}
