/* 
1. Un booleano es un TIPO binario el cual puede tener dos posibles valores “verdadero” y
“falso”
2. Cuando usas un operador de comparación de igualdad, este es una expresión el cual
evalúa a un valor booleano
	a. ==
	b. <=
	c. >=
	d. !=
	e. <
	f. >
*/

package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 7
	b := 42
	fmt.Println(a == b) // Un bool puede ser resultado de una comparacion
}
