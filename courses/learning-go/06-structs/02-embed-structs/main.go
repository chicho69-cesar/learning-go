/* 
Podemos tomar un struct y embeberlo en otro struct. Cuando haces esto el tipo interno es
promovido al tipo externo.
*/

package main

import "fmt"

type persona struct {
	name string
	lastname string
	age int
}

type agenteSecreto struct {
	persona // embedemos un struct dentro de otro struct
	lpm bool
}

func main() {
	p1 := persona {
		name: "Cesar",
		lastname: "Villalobos Olmos",
		age: 21,
	}

	p2 := persona {
		name: "Liz",
		lastname: "Sandoval Vallejo",
		age: 20,
	}

	as1 := agenteSecreto {
		persona: persona { // definimos el valor para el struct embedido
			name: "James",
			lastname: "Bond",
			age: 31,
		},
		lpm: true,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(as1)

	fmt.Printf(
		"Nombre: %s\tApellidos: %s\tEdad: %d\n", 
		p1.name, 
		p1.lastname, 
		p1.age,
	)

	fmt.Printf(
		"Nombre: %s\tApellidos: %s\tEdad: %d\n", 
		p2.name, 
		p2.lastname, 
		p2.age,
	)

	fmt.Printf(
		"Nombre: %s\tApellidos: %s\tEdad: %d\tLicencia para matar: %t\n", 
		as1.name, 
		as1.lastname, 
		as1.age,
		as1.lpm,
	)
}
