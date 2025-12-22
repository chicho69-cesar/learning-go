/*
Este ejercicio te ayudará a reforzar el concepto de method sets:
	Crea un tipo struct persona
	Adjunta el método hablar usando un receptor de tipo puntero
		.*persona
	Crea un tipo interfaz humano
		Para implícitamente implementar esa interfaz, un tipo humano debe tener el
		método hablar
	Crea la función “diAlgo”
		Haz que tome un humano como parámetro
		Haz que llame al método hablar
	Muestra lo siguiente en tu código
		PUEDES pasar un valor de tipo *persona a diAlgo
		NO puedes pasar un valor de tipo persona a diAlgo
	Aquí hay una pista si necesitas ayuda
		https://play.golang.org/p/JQd6LsU_L-K
*/

package main

import "fmt"

type person struct {
	name string
	lastname string
	age int
}

func (p *person) speak() {
	fmt.Println(p.name, "esta hablando")
}

type human interface {
	speak()
}

func tellSomething(h human) {
	h.speak()
}

func main() {
	p := person {
		name: "Cesar",
		lastname: "Villalobos Olmos",
		age: 21,
	}

	tellSomething(&p)
}
