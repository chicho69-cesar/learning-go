package main

import (
	"fmt"
)

func main() {
	n := "Pera"
	switch n { // Evaluamos el valor de n con los cases
		case "Manzana":
			fmt.Println("Manzana Roja")
		case "Pera": // Aqui hara match
			fmt.Println("Pera Verde")
		case "Q":
			fmt.Println("Esta es la q")
		default:
			fmt.Println("Esta es default")
	}
}