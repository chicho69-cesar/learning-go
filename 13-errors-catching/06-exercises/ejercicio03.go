/*
Crea un struct “errorPer” el cual implemente la interfaz builtin.error. Crea una función “foo” que
tenga un valor de tipo error como parámetro. Crea un valor de tipo “errorPer” y pásalo a “foo”.
Si necesitas un pista, aquí hay una.
*/

package main

import (
	"errors"
	"fmt"
	"log"
)

type errorPer struct {
	name string
	lastname string
	err error
}

func (e errorPer) Error() string {
	return fmt.Sprintf("Error en persona: %s %s %s", e.name, e.lastname, e.err)
}

func main() {
	err := errorPer {
		name: "John",
        lastname: "Doe",
        err: errors.New("Error en esta persona"),
	}

	foo(err)
}

func foo(err error) {
	log.Println("Error: ", err, " -> ", err.(errorPer).name)
}
