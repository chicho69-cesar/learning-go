/* 
Usando el código del ejemplo anterior, agrega un registro a tu mapa, ahora imprime el mapa
usando “range”
*/

package main

import "fmt"

func main() {
	persons := map[string][]string {
		"cesar_villalobos": { "Programar", "Futbol", "Peliculas" },
		"liz_sandoval": { "Actuar", "Basketbol", "Musica" },
		"alonso_villalobos": { "Jugar", "Futbol", "Fiestas" },
	}

	persons["hector_garcia"] = []string {
		"Bromear",
		"Lucha Libre",
		"Cine",
		"Musica",
	}

	for k, v := range persons {
		fmt.Println("\n", k)

		for i, val := range v {
			fmt.Printf("\t%d - %s\n", i, val)
		}
	}
}
