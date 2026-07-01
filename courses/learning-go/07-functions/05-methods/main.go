/* 
Un método no es más que una FUNC adjuntada a un TIPO. Cuando adjuntas una función a
un tipo es un método de ese tipo. Adjuntas una función a un tipo con el RECEPTOR.
*/

package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona
	lpm bool
}

// func (r receptor) identificador(parámetros) (retorno(s)) { código }
// agregamos un metodo al tipo agenteSecreto
func (s agenteSecreto) presentar() {
	fmt.Println("Hola, soy", s.nombre, s.apellido)
}

func main() {
	sa1 := agenteSecreto {
		persona: persona {
			"Cesar",
			"Villalobos Olmos",
		},
		lpm: true,
	}

	sa2 := agenteSecreto {
		persona: persona {
			"Liz",
			"Sandoval Vallejo",
		},
		lpm: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa2)

	sa1.presentar()
	sa2.presentar()
}
