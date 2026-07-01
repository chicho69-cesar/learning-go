/* 
- Un simple valor que no cambia.
Sólo existe en el momento de compilación..
Hay constantes con TIPO y SIN TIPO
	const hola = "Hola, mundo"
	const typedHello string = "Hello, World"
Constante SIN TIPO
	Un valor constante que no tiene un tipo fijo
		- “constante de algún tipo”
		- No es forzada a obedecer las reglas estrictas que previenen
		combinar diferentemente valores con un tipo.
	Una constante sin tipo puede ser implícitamente convertida por el compilador.
Es este concepto de constante sin tipo lo que hace posible que usemos
constantes en Go con libertad.
	Muy útil, por ejemplo
		Cuál es el tipo de 42?
			int?
			uint?
			float64?
		Si no tuviéramos constante SIN TIPO (constantes de algún tipo),
		tendríamos que hacer conversión en cada valor literal que usamos.
			Y eso no sería muy agradable
*/

package main

import (
	"fmt"
)

const a = 42 // Definimos una constante
const b = 42.32
const c = "Cesar Villalobos Olmos" // Constante con tipo inferido
const d string = "chicho69-cesar" // Constante con tipo definido

type nombre string
var e nombre

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

	/* Aqui le asigamos el valor a una variable el valor de una constante, 
	y aunque estemos creando una variable, el tipo adyacente de esta es una
	constante, por lo que esta va a tomar como valor inicial el valor de 
	la constante a la que le definimos */
	e = c
	fmt.Println(e)
}
