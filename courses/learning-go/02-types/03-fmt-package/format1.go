/*
Equivalencias en el defalut format para imprimir con Printf y el 
caracter de estilo %v, estos son conocidos como VERBOS: 
	bool:                    %t
	int, int8 etc.:          %d
	uint, uint8 etc.:        %d, %#x if printed with %#v
	float32, complex64, etc: %g
	string:                  %s
	chan:                    %p
	pointer:                 %p
*/

package main

import (
	"fmt"
)

var a int
var b string = "Programa"

func main() {
	fmt.Println(a)
	fmt.Println(b)

	// Escribimos con formato
	fmt.Printf("El valor de la variable a es: %d\n", a)
	fmt.Printf("El valor de la variable b es: %s", b)
}
