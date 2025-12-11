/* 
Usando un COMPOSITE LITERAL:
Crea un SLICE de TIPO int
Asigna 10 VALORES
Haz Range sobre el slice e imprime los valores.
Usando format para imprimir
	Imprime el TIPO del slice
*/

package main

import "fmt"

func main() {
	slice := []int{ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 }
	fmt.Printf("El tipo es: %T\n\n", slice)

	for i, v := range slice {
		fmt.Printf("%d - %d\n", i, v)
	}
}
