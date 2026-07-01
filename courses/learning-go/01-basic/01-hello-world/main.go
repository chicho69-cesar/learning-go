package main

import "fmt"

func main() {
	/* Aqui el Println dos devuelve tambien el numero de 
	bytes y el error */
	nb, err := fmt.Println("Hola Mundo!!!")
	fmt.Println(nb)
	fmt.Println(err)
}
