/* 
Usando un COMPOSITE LITERAL:
Crea un ARREGLO el cual tenga 5 VALORES de TIPO int
Asigna VALORES a cada posición dada por los índices.
Itera con Range sobre el arreglo e imprime los valores.
Usando el paquete fmt
	Imprime el TIPO del arreglo
*/

package main

import "fmt"

func main() {
	var array [5]int
	fmt.Println(array)
	fmt.Printf("%T\n", array)

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5
	fmt.Println("\n", array)

	fmt.Println()
	for i, v := range array {
		fmt.Printf("%d - %d\n", i, v)
	}
}
