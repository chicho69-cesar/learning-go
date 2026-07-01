/* 
Bit shifting es cuando desplazan dígitos binarios a la izquierda o a la derecha. Podemos
usar bit shifting junto con iota, para construir constantes creativas.

Operados de Bit Shifting
a >> b: Desplaza b digitos binarios a la deracha el valor
binario de a
a << b: Desplaza b digitos binarios a la izquierda el valor
binario de a

Por ejemplo: 
decimal 		binario
4				100
8 				1000
6 				110
12 				1100

Si en los numeros anterios hacemos lo siguiente: 
	4 << 1 -> Obtenemos -> 1000 que es 8
	4 >> 1 -> Obtenemos -> 10 que es 2
	6 << 1 -> Obtenemos -> 1100 que es 12

Con el verbo %b conviertes un numero en su version binaria
*/

package main

import (
	"fmt"
)

func main() {
	x := 6
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	y = x >> 1
	fmt.Printf("%d\t\t%b\n", y, y)
}
