/* 
Un slice multidimensional es como una hoja de cálculo. Puede tener un slice de slices de
algún tipo. Suena confuso? En este video lo aclaramos.
*/

package main

import (
	"fmt"
)

func main() {
	// Creamos un slice
	cv := []string{ "Cesar", "Villalobos", "Soccer", "Basketball", "Volleyball" }
	fmt.Println(cv)

	// Creamos otro slice
	lz := []string{ "Liz", "Sandoval", "Soccer", "Beisball", "Tenis" }
	fmt.Println(lz)

	// Creamos un slice el cual va a contener slices dentro, es decir, una matriz
	vp := [][]string{ cv, lz }
	fmt.Println(vp)
}
