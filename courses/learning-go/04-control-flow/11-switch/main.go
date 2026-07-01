/* 
Sintaxis de un Switch: 
	switch expresion {
		case opcion1: 
			Accion si expresion == opcion1
		case opcion2: 
			Accion si expresion == opcion2
		...
		default: 
			Accion si ninguna opcion es igual a la expresion
	}
*/

package main

import (
	"fmt"
)

func main() {
	/* Cuando no usamos una expresion en el switch esta se toma como verdadera
	asi que una vez que un case haga match con true se ejecutara y alli
	se ejecutaran un break aunque este no se defina, por lo que los demas 
	cases ya no se ejecutaran */
	switch {
		case false:
			fmt.Println("Esta no se debe imprimir")
		case (2 == 4):
			fmt.Println("Esta no se debe imprimir2")
		case (3 == 3):
			fmt.Println("Imprime")
		case (4 == 4):
		fmt.Println("también verdadera, se imprime?")
	}
}
