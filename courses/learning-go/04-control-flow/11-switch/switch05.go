package main

import (
	"fmt"
)

func main() {
	n := "Pera"
	switch n {
		/* Podemos hacer multiples evaluaciones en un solo case y con 
		que una de las opciones haga match se tomara como true*/
		case "Manzana", "Pera", "Mango":
			fmt.Println("Varias frutas")
		case "M":
			fmt.Println("m")
		case "Q":
			fmt.Println("Esta es la q")
		default:
			fmt.Println("Esta es default")
	}
}
