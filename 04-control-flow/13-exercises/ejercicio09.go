/* 
Crea un programa que use una declaración switch con la expresión de switch especificada
como una variable de TIPO string y el IDENTIFICADOR “deporteFav”.
*/

package main

import "fmt"

func main() {
	deporteFav := "Futbol"

	switch deporteFav {
		case "Basquetbol":
			fmt.Println("Su deporte favorito es el: Basquetbol")
		case "Tenis":
			fmt.Println("Su deporte favorito es el: Tenis")
		case "Voleybol":
			fmt.Println("Su deporte favorito es el: Voleybol")
		case "Futbol":
			fmt.Println("Su deporte favorito es el: Futbol")
		case "Beisbol":
			fmt.Println("Su deporte favorito es el: Beisbol")
		default:
			fmt.Println("No se cual sea su deporte favorito :(")
	}
}
