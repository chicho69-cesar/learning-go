/*
Una declaración "defer" (diferir) invoca a una función cuya ejecución es diferida para
el momento en el que la función donde está contenida retorna, en cualquiera de los
siguientes casos: o la función ejecuta una declaración de retorno, o llega al final de su
cuerpo o porque la gorutina (goroutine) correspondiente está en pánico (panicking)
*/

package main

import "fmt"

func main() {
	/* Aqui se va a ejecutar primero todo el codigo de la función, en
	este caso main, y una vez que retorne, llega al final de su cuerpo
	o la goroutine correspondiente entra en panico */

	defer foo()
	bar() 
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
