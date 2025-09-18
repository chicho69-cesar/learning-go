package main

import (
	"fmt"
)

var z = 21 // Scope a nivel de paquete

func main() {
	/* Esto es igual a poner: 
	var x int = 42 + 7 */
	x := 42 + 7
	/* Scope a nivel de funcion */
	y := "Cesar Villalobos Olmos"

	fmt.Println(x)
	fmt.Println(y)

	x = 50

	fmt.Println(x)
	fmt.Println(z)
}

func numero() {
	fmt.Println(z)
}
