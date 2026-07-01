package main

import (
	"fmt"
)

func main() {
	/* Cuando en un switch usamos fallthrough es como usar un anti break
	ya que el switch seguira evaluando los demas cases aunque uno ya 
	haya echo match con true */
	switch {
		case false:
			fmt.Println("Esta no se debe imprimir")
		case (2 == 4):
			fmt.Println("Esta no se debe imprimir2")
		case (3 == 3):
			fmt.Println("Imprime")
			fallthrough
		case (4 == 4):
			fmt.Println("también verdadera, se imprime?")
	}
}
