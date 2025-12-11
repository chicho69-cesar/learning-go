/* 
Crea y usa un struct anónimo.
*/

package main

import "fmt"

func main() {
	me := struct {
		name string
		lastname string 
		age int 
		friends map[string]int 
		drinks []string
	} {
		name: "Cesar",
		lastname: "Villalobos Olmos",
		age: 21,
		friends: map[string]int {
			"Liz": 1,
			"Alonso": 2,
			"Daniel": 3,
		},
		drinks: []string {
			"Cafe",
			"Cerveza",
			"Tequila",
		},
	}

	fmt.Println(me.name, me.lastname, me.age)
	fmt.Println("AMIGOS")

	for k, v := range me.friends {
		fmt.Printf("\t%d: %s\n", v, k)
	}

	fmt.Println("BEBIDAS")

	for i, v := range me.drinks {
		fmt.Printf("\t%d: %s\n", i + 1, v)
	}
}
