package main

import (
	"fmt"
)

func main() {
	switch {
		case false:
			fmt.Println("Esta no se debe imprimir")
		case (2 == 4):
			fmt.Println("Esta no se debe imprimir2")
		default:
			fmt.Println("Este es default si se imprime")
	}
}
