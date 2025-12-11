/* 
UN SLICE almacena VALORES del mismo TIPO. Si quisiéramos almacenar todos nuestros
números favoritos usamos un SLICE de ints. Si quisiera almacenar todas mis comidas favoritas
usaría un SLICE de strings. Usaremos un LITERAL COMPUESTO para crear un slice. Un
literal compuesto es cuando colocamos el TIPO seguido de LLAVES y luego colocamos los
valores en el área dentro de las llaves.
*/

package main

import (
	"fmt"
)

func main() {
	//tipo{elementos} //COMPOSITE LITERAL
	/* Los slices tienen un tamaño variado, a diferencia de los arrays
	que tienen un tamaño fijo */
	x := []int{ 1, 2, 3, 4, 5 } // creamos un slice de datos enteros
	fmt.Println(x)
	fmt.Printf("%T", x)
}

//Usamos slices para agrupar VALORES del mismo TIPO