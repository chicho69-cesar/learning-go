/* 
1. A nivel de paquete usando la palabra clave “var”, crear una VARIABLE con el
IDENTIFICADOR “y”. La variable debería ser del mismo TIPO SUBYACENTE que tu
TIPO “x” creado anteriormente
	a. ejemplo:
2. en func main
	a. Esto lo debería tener listo
		i. Imprimir el valor de la variable “x”
		ii. Imprimir el tipo de la variable “x”
		iii. Asigna un VALOR a la VARIABLE “x” usando el OPERADOR “=”
		iv. Imprime el valor de la variable “x”
	b. Ahora haces esto
		i. Ahora usa CONVERSIÓN para convertir el TIPO del VALOR almacenado
		en “x” al TIPO IMPLÍCITO
			1. Usa el operador “=” para ASIGNAR ese valor a “y”
			2. Imprime el valor almacenado en “y”
			3. Imprime el tipo de “y”
*/

package main

import "fmt"

type myType int

var x myType
var y int

func main() {
	fmt.Printf("El valor de x es: %v\n", x)
	fmt.Printf("El tipo de x es: %T\n", x)
	x = 42
	fmt.Printf("El nuevo valor de x es: %v\n", x)

	y = int(x)
	fmt.Printf("El valor de y es: %v\n", y)
	fmt.Printf("El tipo de y es: %T\n", y)
}
