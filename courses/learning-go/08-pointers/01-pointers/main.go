/*
Que son punteros (pointers)?
.Todos los valores son almacenados en memoria. Cada ubicación en memoria tiene una
dirección. Un pointer es una dirección de memoria.
&
*int
*
*/

package main

import "fmt"

func main() {
	a := 21
	fmt.Println(a)
	fmt.Println(&a) // accedemos a la direccion en memoria donde se almacena la variable a

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // mostramos el tipo de la direccion en memoria

	b := &a // le asignamos a b la direccion de memoria de a
	fmt.Println(b) // es igual a poner &a
	fmt.Println(*b) /* Esto es un puntero, con el * accedemos al valor que 
	esta almacenado en la direccion de memoria especificada */
	fmt.Println(*&b) // accedemos al valor almacenado en la direccion de memoria de b

	*b = 20 // esto es como poner a = 20
	fmt.Println(a)
}
