/* 
Puedes retornar una función desde una función.
*/

package main

import (
	"fmt"
)

func main() {
	s1 := foo() // asignamos el valor de retorno
	fmt.Println(s1)

	x := bar() // función que regresa una función

	fmt.Printf("%T\n", x)

	i := x() // función que fue devuelta
	fmt.Println(i)

	fmt.Println(bar()()) // llamamos directamente a la función que nos es retornada
}

func foo() string {
	s := "Hola mundo."
	return s
}

// clousure
func bar() func() int {
	return func() int {
		return 123
	}
}
