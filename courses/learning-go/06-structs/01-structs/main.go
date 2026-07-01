/* 
Un Struct es un tipo de dato compuesto. (tipos de datos compuestos, también, tipos de datos
agregados o tipos de datos complejos). Los structs nos permiten componer valores con
otros valores de diferentes tipos.
*/

package main

import "fmt"

/* Creamos una estructura persona */
type persona struct {
	name string
	lastname string
	age int
}

func main() {
	// creamos una instancia de una estructura
	p1 := persona {
		name: "Cesar",
		lastname: "Villalobos Olmos",
		age: 21,
	}

	p2 := persona {
		name: "Liz",
		lastname: "Sandoval Vallejo",
		age: 20,
	}

	// imprimimos una estructura
	fmt.Println(p1)
	fmt.Println(p2)

	// Imprimimos los campos por separado de la estructura
	fmt.Printf("Nombre: %s\tApellidos: %s\tEdad: %d\n", p1.name, p1.lastname, p1.age)
	fmt.Printf("Nombre: %s\tApellidos: %s\tEdad: %d\n", p2.name, p2.lastname, p2.age)
}
