/*
Podemos agregar información a nuestros errores con
● errors.New()
	fmt.Errorf()
● builtin.error
“Los valores Error en Go no son especiales, son como cualquier otro valor, y por ende tienes el
lenguaje entero a tu disposición.” - Rob Pike
*/

package main

import (
	"errors"
	"log"
	"math"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// esto es como lanzar un error con throw en otros lenguajes
		return 0, errors.New("De matemática elemental: no hay raíz cuadrada real de un número negativo")
	}

	return math.Sqrt(f), nil
}
