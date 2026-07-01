/* 
También podemos crear structs anónimos. Un struct anónimo es un struct el cual no es
asociado con un identificador en específico.
*/

package main

import "fmt"

// estructura definida
type persona struct {
	name string
	last string
	age int
}

func main() {
	// creamos una instancia de la estructura persona
	p1 := persona {
		name: "Cesar",
		last: "Villalobos Olmos",
		age: 21,
	}

	fmt.Println(p1)

	/* Creamos una struct anonima, es decir, esta no esta definida
	en ningun tipo ni identificador, por lo que solo lo podremos usar
	aqui, y no podremos crear mas instancias de esta struct, a menos
	que volvamos a escribir todo el mismo codigo */
	p2 := struct {
		name, last string // podemos declarar todas las props del mismo tipo
		age int
	} {
		name: "Liz",
		last: "Sandoval Vallejo",
		age: 20,
	}

	fmt.Println(p2)
}
