/*
Crea un strutc person
Crea una función llamada “cambiame” con un *person como parámetro
	Cambia el valor almacenado en la dirección de memoria del *person
		Importante
			Para desreferenciar un struct, usa (*valor).campo
				me.name
				(*me).name
			“Como una excepción, si el tipo de x es un tipo puntero con
			name y (*x).c es una expresión válida de selección denotando
			un campo (pero no un método), x.c es una forma abreviada de
			(*x).c”.
				https://golang.org/ref/spec#Selectors
Crea un valor de tipo person
	Imprime el valor
En func main
	llama “cambiame”
En func main
	Imprime el valor
*/

package main

import "fmt"

type person struct {
	name string
}

func main() {
	me := person {
		name: "Cesar Villalobos Olmos",
	}

	fmt.Println(me)
	changeMe(&me, "Liz Sandoval Vallejo")
	fmt.Println(me)
}

func changeMe(p *person, newName string) {
	/* Ambas funciones funcionan correctamente */
	p.name = newName
	(*p).name = newName
}
