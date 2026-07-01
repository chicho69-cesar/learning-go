package main

import (
	"fmt"
)

//* Es una mala practica tener muchas variables en scope a nivel de package

// var z int = 21 Le podemos asignar el valor directamente
var z int // La podemos solo declarar

func main() {
	z = 21
	fmt.Println(z)
}
