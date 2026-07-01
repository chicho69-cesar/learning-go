/*
Asignando una func a una variable
*/

package main

import "fmt"

func main() {
	// le podemos asignar una función a una variable
	f := func() {
		fmt.Println("Expresión función")
	}
	f() // ejecutamos la función expresión

	// función expresion que recibe un parámetro
	g := func(x int) {
		fmt.Println("El año del descubrimiento de América fue:", x)
	}
	g(1492)

	// función expresion que regresa un valor
	r := func (x int) int {
		return x * 2
	}
	fmt.Println("El doble de 15 es = ", r(15))
}
