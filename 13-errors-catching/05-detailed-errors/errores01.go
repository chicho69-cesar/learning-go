package main

import (
	"fmt"
	"log"
	"math"
)

var ErrMath error // creamos nuestro tipo error

func main() {
	_, err := sqrt(-10.56)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		ErrMath = fmt.Errorf("De matemática elemental: no hay raíz cuadrada real de un número negativo: %v", x)
		return 0, ErrMath
	}

	return math.Sqrt(x), nil
}
