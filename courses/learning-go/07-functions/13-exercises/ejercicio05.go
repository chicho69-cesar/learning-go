/*
# Crea un tipo CUADRADO
# Crea un tipo CÍRCULO
# Adjunta un método a cada uno que calcule y retorne el ÁREA
	Área de un círculo= π r 2
	Área de un cuadrado = L * H
# Crea un tipo FORMA que defina una interfaz como cualquier cosa que tenga el método
ÁREA
# Crea una func INFO el cuál toma un tipo forma e imprime el área de la misma.
# Crea un valor de tipo cuadrado
# Crea un valor de tipo círculo
# Usa la func info para imprimir el área de un cuadrado
# Usa la func info para imprimir el área de un círculo
*/

package main

import (
	"fmt"
	"math"
)

type square struct {
	large float64
	height float64
}

func (s square) area() float64 {
	return s.large * s.height
}

type circle struct {
	radio float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radio, 2)
}

type form interface {
	area() float64
}

func getArea(f form) {
	fmt.Printf("El area es: %g\n", f.area())
}

func main() {
	sq := square {
		large: 10.0,
		height: 5.0,
	}

	ci := circle {
		radio: 2,
	}

	getArea(sq)
	getArea(ci)
}
