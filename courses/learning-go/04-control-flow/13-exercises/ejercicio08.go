/* 
Crea un programa que use una declaración switch sin expresión especificada.
*/

package main

import "fmt"

func main() {
	hero := "Batman"

	switch {
		case hero == "Super Man":
			fmt.Println("El heroe es Super Man")
		case hero == "Aqua Man":
			fmt.Println("El heroe es Aqua Man")
		case hero == "Batman":
			fmt.Println("El heroe es Batman")
		case hero == "Wonder woman":
			fmt.Println("El heroe es Wonder woman")
		case hero == "Flash":
			fmt.Println("El heroe es Flash")
		default:
			fmt.Println("El heroe no esta en la Justice League")
	}
}
