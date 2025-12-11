package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	/* Aunque no volvamos a asignar el slice a la variable x, como un slice
	es una referencia o puntero a una direccion de memoria donde se guarda
	el arreglo, esta función si muta el arreglo aunque no lo volvamos a asignar */
	_ = append(x[:2], x[5:]...) // the same underlying array stores the value of the new slice

	fmt.Println(x)
}
