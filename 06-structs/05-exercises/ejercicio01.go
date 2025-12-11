/*
Crea tu propio tipo “persona” el cual tendrá un tipo subyacente tipo “struct” de manera que
pueda almacenar la siguiente data:
	Nombre
	Apellido
	Sabores de helado favoritos
Crea dos VALORES de TIPO persona. Imprime los valores, usa range sobre los elementos en
el slice el cual almacena los valores de helados favoritos.
*/

package main

import "fmt"

type person struct {
	name string
	lastname string
	icecreams []string
}

func main() {
	p1 := person {
		name: "Cesar",
		lastname: "Villalobos Olmos",
		icecreams: []string {
			"Strawberry",
			"Pineapple",
			"Napolitano",
		},
	}

	p2 := person {
		name: "Liz",
		lastname: "Sandoval Vallejo",
		icecreams: []string {
			"Gansito",
			"Napolitano",
			"Chocolate",
		},
	}

	fmt.Println("Datos de la persona 1: ")
	fmt.Printf("Nombre: %s\nApellidos: %s\nHelados: \n", p1.name, p1.lastname)

	for i, v := range p1.icecreams {
		fmt.Printf("\t%d: %s\n", i, v)
	}

	fmt.Println()

	fmt.Println("Datos de la persona 2: ")
	fmt.Printf("Nombre: %s\nApellidos: %s\nHelados: \n", p2.name, p2.lastname)

	for i, v := range p1.icecreams {
		fmt.Printf("\t%d: %s\n", i, v)
	}
}
