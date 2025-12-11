/* 
Crea un mapa con una llave TIPO string el cual representa el “nombre_apellido” de una
persona y un valor de TIPO []string el cual almacena sus cosas favoritas. Almacena tres
registros en tu mapa. Imprime todos sus valores con su índice de posición en el slice.
	`eduar_tua`, `computadoras`, `montaña`, `playa`
	`carlos_ramon`, `leer`, `comprar`, `música`
	`juan_bimba`, `helado`, `pintar`, `bailar`
*/

package main

import "fmt"

func main() {
	persons := map[string][]string {
		"cesar_villalobos": { "Programar", "Futbol", "Peliculas" },
		"liz_sandoval": { "Actuar", "Basketbol", "Musica" },
		"alonso_villalobos": { "Jugar", "Futbol", "Fiestas" },
	}

	for k, v := range persons {
		fmt.Println("\n", k)

		for i, val := range v {
			fmt.Printf("\t%d - %s\n", i, val)
		}
	}
}
