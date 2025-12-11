/* 
Ejercicio: Mostrar los numeros pares hasta un numero dado por el usuario
*/

package main

import "fmt"

var limit int

func main() {
	fmt.Print("Escribe el limite de los numeros: ")
	fmt.Scan(&limit)

	fmt.Println("\nNumeros pares: ")
	for i := 1; i <= limit; i++ {
		if i % 2 == 0 {
			fmt.Printf("\t%d\n", i)
		}
	}
}
