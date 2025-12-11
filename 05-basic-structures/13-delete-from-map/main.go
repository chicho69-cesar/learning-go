/* 
Borras un elemento de un mapa usando delete(<nombre del mapa>, “llave”). No arroja
ningún error si usas usa llave que no existe. Para confirmar que borraste un par llave-valor,
verifica que el par existe con el idioma coma ok.
*/

package main

import "fmt"

func main() {
	m := map[string]int {
		"Batman": 32,
		"Robin":  27,
	}

	fmt.Println(m)

	delete(m, "Batman")
	fmt.Println(m)

	if v, ok := m["Robin"]; ok {
		fmt.Println("Eliminando la llave con el valor", v)
		delete(m, "Robin") // eliminamos el elemento con la clace Robin del mapa m
	}

	fmt.Println(m)
}
