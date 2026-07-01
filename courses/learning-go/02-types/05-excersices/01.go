/* 
1. Usando el operador de declaración corta, ASIGNA los siguientes VALORES a
VARIABLES con los IDENTIFICADORES “x”, “y” y “z”
	a. 42
	b. “James Bond”
	c. true
2. Luego imprime los valores almacenados en esas variables usando
	a. Una sola declaración de la función println
	b. Múltiples declaraciones de println
*/

package main 

import "fmt"

func main() {
	x := 42
	y := "James Bond”"
	z := true

	fmt.Println(x, y, z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
