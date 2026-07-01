/* 
Crear un slice de slice de string ([][]string). Almacena los siguientes valores en un slice
multi-dimensional:
	"Batman", "Jefe", "Vestido de negro"
	"Robin", "Ayudante", "Vestido de colores"
Haz range sobre los registros, luego sobre los datos de cada registro.
*/

package main

import "fmt"

func main() {
	heroes := [][]string {
		{ "Batman", "Jefe", "Vestido de negro" },
		{ "Robin", "Ayudante", "Vestido de colores" },
	}

	for i, heroe := range heroes {
		fmt.Println("Heroe numero ", i + 1)

		for _, value := range heroe {
			fmt.Println("\t", value)
		}
	}
}
