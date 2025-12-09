/* 
Enteros
	Números sin decimales
		También conocidos como enteros
	int & uint
		“Tamaños de implementación-específica”
	xTodos los tipos numéricos son diferentes excepto
		byte el cual es un alias para uint8
		Rune el cual es un alias para int32
	Los tipos son únicos
		- “Esto es un lenguaje de programación estático”
		- “Las conversiones son requeridas cuando mezclamos diferentes tipos de
		datos numéricos en una expresión o asignación. Por ejemplo, int32 e int
		no son del mismo tipo aún cuando pueden tener el mismo tamaño en una
		arquitectura particular.” (fuente)
	Regla de oro: solo usa int
floating point
	Números con decimales
		También conocidos como números reales
	Regla de oro: solo usa float64
Overflows
*/

package main

import "fmt"

var a int
var b int8
var c int16
var d int32
var e int64

var f float32
var g float64

func main() {
	//* TIPOS ENTEROS
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	//* TIPOS FLOTANTES
	fmt.Println(f)
	fmt.Println(g)
}
