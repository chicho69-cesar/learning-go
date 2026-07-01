/* 
Crea un ciclo for usando esta sintaxis
	for { }
Haz que imprima los años que has vivido
*/

package main

import "fmt"

func main() {
	myAge, counter := 20, 0

	for {
		fmt.Printf("Año: %d\n", counter + 2002)
		counter++

		if myAge < counter {
			break
		}
	}
}
