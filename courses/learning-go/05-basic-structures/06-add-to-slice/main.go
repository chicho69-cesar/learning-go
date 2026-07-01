/* 
Para añadir valores a un slice, debemos usar la función especial integrada append. Esta
función retorna un slice del mismo tipo.
*/

package main

import "fmt"

func main() {
	x := []int{ 4, 5, 7, 8, 42 } // creamos el slice
	fmt.Println(x) // lo imprimimos
	x = append(x, 44, 55, 66) /* le agregamos datos al final del slice x
	este metodo regresa un slice con los elementos agregados */
	fmt.Println(x)

	y := []int{ 333, 444, 666 } // creamos otro slice
	/* El operador ... es como el spread operator en JavaScript, ya que 
	desestructura y nos devuelve todos los elementos del slice */
	x = append(x, y...)
	fmt.Println(x)
}
