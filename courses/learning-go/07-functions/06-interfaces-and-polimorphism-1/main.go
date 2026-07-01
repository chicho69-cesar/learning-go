/*
En Go, los valores pueden ser de más de un tipo. Una interfaz permite a un valor ser de más
de un tipo. Creamos una interfaz usando la sintaxis : “palabra clave identificador tipo” entonces
para una interfaz sería: “type humano interface” Luego definimos cuáles método(s) debe tener
un tipo para implementar esa interfaz. Si un TIPO tiene los métodos requeridos, el cual podrían
ser ninguno, (la interfaz vacía denotada por interface{}), entonces ese TIPO implícitamente
implementa la interfaz y es también de ese tipo de interfaz. En Go, los valores pueden ser más
de un tipo.
*/

package main

import "fmt"

// creamos una estructura person
type person struct {
	name string
	lastanme string
}

// creamos una estructura agente secreto que tiene un tipo embedido person
type secretAgent struct {
	person
	lpm bool
}

// creamos un metodo para las personas
func (p person) presentate() {
	fmt.Println("Hola soy ", p.name, p.lastanme, " - La persona se presenta")
}

// creamos polimorfismo para el metodo presentate en la struct secretAgent
func (s secretAgent) presentate() {
	fmt.Println("Hola soy ", s.name, s.lastanme, " - El agente secreto se presenta")
}

/* Creamos una interfaz y los tipos que implementen todos los metodos que
posee la interfaz, estaran implementando dicha interfaz.
Por lo tanto aqui los tipos person y secretAgent (structs) implementan
la interfaz human */
type human interface {
	presentate()
}

func show(h human) {
	fmt.Println("\nPresentacion desde interfaz humano")
	h.presentate()
}

func main() {
	sa1 := secretAgent {
		person: person {
			name: "Cesar",
			lastanme: "Villalobos Olmos",
		},
		lpm: false,
	}

	sa2 := secretAgent {
		person: person {
			name: "Liz",
			lastanme: "Sandoval Vallejo",
		},
		lpm: true,
	}

	p := person {
		name: "Danna",
		lastanme: "Delgado Hernandez",
	}

	fmt.Println(sa1)

	sa1.presentate()
	sa2.presentate()
	p.presentate()

	show(sa1)
	show(sa2)
	show(p)
}
