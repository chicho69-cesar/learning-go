/* 
Usa el código del ejemplo anterior y almacena los valores de tipo persona en un mapa con la
llave apellido. Accede a cada valor en el mapa. Imprime los valores. Imprime también los
valores haciendo range sobre el slice
*/

package main

import "fmt"

type person struct {
	name string
	lastname string
	icecreams []string
}

func main() {
	persons := map[string]person {}

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

	persons[p1.lastname] = p1
	persons[p2.lastname] = p2

	fmt.Println("Datos de la personas: ")

	for _, p := range persons {
		fmt.Printf("Nombre: %s\nApellidos: %s\nHelados: \n", p.name, p.lastname)

		for i, v := range p.icecreams {
			fmt.Printf("\t%d: %s\n", i, v)
		}

		fmt.Println()
	}
}
