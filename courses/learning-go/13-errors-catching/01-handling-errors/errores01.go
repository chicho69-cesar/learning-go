package main

import "fmt"

func main() {
	/* Todas las funciones que implementan la interfaz error, nos 
	suelen devolver en la lista de retornos un error por si es que
	algo inesperado sucede, sino regresa nil */
	n, err := fmt.Println("Hello world!")
	if err != nil { // si ocurrio un error
        fmt.Println(err)
    }
	fmt.Println(n)
}
