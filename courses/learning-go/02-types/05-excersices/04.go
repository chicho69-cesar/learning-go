/* 
1. Crea tu propio tipo. Haz que el tipo subyacente, raíz o implícito sea un int.
2. Crea una VARIABLE de tu nuevo TIPO con el IDENTIFICADOR “x” usando la palabra
clave “VAR”
3. En func main
	a. Imprime el valor de la variable “x”
	b. Imprime el tipo de la variable “x”
	c. Asigna 42 a la VARIABLE “x” usando el OPERADOR “=”
	d. Imprime el valor de la variable “x”
*/

package main

import "fmt"

type myType int

var x myType

func main() {
	fmt.Printf("El valor es: %v\n", x)
	fmt.Printf("El tipo es: %T\n", x)

	x = 42

	fmt.Printf("El nuevo valor es: %v\n", x)
}
