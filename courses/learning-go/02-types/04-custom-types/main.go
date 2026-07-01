package main

import "fmt"

/* Creamos nuestro nuevo tipo de datos, el cual es un tipo
subyacente de int */
type dinero int 

var a int
var b dinero

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 1000
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	
	/* Aunque dinero sea un tipo de dinero que es subyacente de int
	no podemos asignar esto a un int, debemos de hacerle uncasting,
	lo que nos va a convertir dinero a int */
	a = int(b) 
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
