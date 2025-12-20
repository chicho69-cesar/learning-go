//* RECEPTOR NO-POINTER Y VALOR POINTER

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

//* receptor = No Pointer
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func info(s form) {
	fmt.Println("Area: ", s.area())
}

func main() {
	c := circle { 5 } // radius = 5
	info(&c) //* valor = Pointer
}
