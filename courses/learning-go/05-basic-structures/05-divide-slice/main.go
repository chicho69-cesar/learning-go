/* 
Podemos dividir un Slice, lo que quiere decir que podemos cortar y desechar partes de un
slice. Hacemos esto con el operador dos puntos ( : ).
*/

package main

import "fmt"

func main() {
	x := []int{ 1, 2, 3, 4, 5 }
	fmt.Println(x)
	fmt.Println(len(x)) // Tamaño del slice
	fmt.Println(cap(x)) // Capacidad del slice o hasta donde puede llegar
	fmt.Println(x[:]) // creamos un slice cortando los elementos de 0 hasta len
	fmt.Println(x[1:3]) // cortamos el slice de la pos 1 hasta antes de la 3
	fmt.Println(x[2:4]) // cortamos el slice de la pos 2 hasta antes de la 4

	for i, v := range x {
		fmt.Println(i, " ", v) // index - valor
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i]) // index - valor
	}
}
