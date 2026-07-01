/* 
● Crea un valor y asígnalo a una variable.
● Imprime la dirección de ese valor.
*/

package main

import "fmt"

func main() {
	me := map[string]string {
		"name": "Cesar",
		"lastname": "Villalobos Olmos",
	}

	fmt.Println(me)
	fmt.Println(&me)
}
