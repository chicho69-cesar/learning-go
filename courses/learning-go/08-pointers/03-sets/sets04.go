//* RECEPTOR POINTER Y VALOR NO-POINTER

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type form interface {
	area() float64
}

//* receptor = Pointer
func (c *circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s form) {
	fmt.Println("Area: ", s.area())
}

func main() {
	c := circle { 5 } // radius = 5
	//! No Funciona tal y como se muestra en la tabla del README
	//info(c) //* valor = No Pointer

	fmt.Println(c.area())
}
