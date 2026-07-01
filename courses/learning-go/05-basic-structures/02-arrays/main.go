/* 
Los arreglos son mayormente usados como un bloque constructor en el lenguaje de
programación Go. En algunas circunstancias, podríamos usar un array, pero en la mayoría de
los casos la recomendación es usar slices en vez de arreglos.
*/

package main

import "fmt"

func main() {
	var x [5]int // definimos un arreglo de enteros de 5 elementos
	fmt.Println(x) // mostramos el arreglo
	fmt.Printf("Tipo: %T\n", x) // mostramos el tipo arreglo

	x[3] = 42 // accedemos a un elemento del arreglo

	fmt.Println(x)
	fmt.Println(len(x)) // mostramos el tamaño del arreglo
}
