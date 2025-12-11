/* 
Crea un ciclo usando la siguiente sintaxis
	for condición { }
Haz que imprima los años que has vivido
*/

package main

import "fmt"

func main() {
	myAge, counter := 20, 0

	for counter <= myAge {
		fmt.Printf("Año: %d\n", counter + 2002)
		counter++
	}
}
