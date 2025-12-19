/*
Crea un struct con
	El identificador “persona”
	Los campos:
		Nombre
		Apellido
		Edad
Adjunta un método al tipo persona con
	El identificador “presentar”
	El método debe hacer que el tipo persona diga su nombre y edad
Crea un valor de tipo persona
Llama al método usando el valor tipo persona
*/

package main

import "fmt"

type person struct {
	name string
	lastname string
	age int
}

func (p person) presentate() {
	fmt.Println("Hola soy ", p.name, p.lastname, " y tengo ", p.age, " años")
}

func main() {
	p := person {
		name: "Cesar",
		lastname: "Villalobos Olmos",
		age: 21,
	}

	p.presentate()
}
