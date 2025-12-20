/* 
el paquete sort tiene funciones que nos permiten ordenar los elementos
de un slice o de una coleccion
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int { 4, 7, 3, 42, 99, 18, 16, 56, 12 }
	strings := []string { "James", "Q", "M", "Moneypenny", "Dr. No" }

	fmt.Println(ints)
	sort.Ints(ints) // ordenamos un slice de ints
	fmt.Println(ints)

	fmt.Println(strings)
	sort.Strings(strings) // ordenamos un slice de strings
	fmt.Println(strings)
}
