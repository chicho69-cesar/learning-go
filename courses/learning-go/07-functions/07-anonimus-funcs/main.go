/*
Funciones anónimas auto-ejecutables
*/

package main

import "fmt"

func main() {
	foo()

	/* Creamos una función anonima que se manda a llamar automaticamente
	desde su definicion */
	func() {
		fmt.Println("La función anonima se auto-ejecuto")
	}()

	/* Creamos una funion anonima que recibe parametros y al autoinvocarse
	le pasamos los argumentos que reciba */
	func(x int) {
		fmt.Println("Funcion anonima con parametros, ", x)
	}(10)
}

func foo() {
	fmt.Println("La función foo se ejecuto")
}
