package main

import "fmt"

func main() {
	switch {
		case false:
			fmt.Println("Esta no se debe imprimir")
		case (2 == 4):
			fmt.Println("Esta tampoco se debe imprimir")
		case (3 == 3):
			fmt.Println("Se imprime")
			fallthrough
		case (4 == 4):
			fmt.Println("Tambien se imprime")
			fallthrough
		case (7 == 9):
			fmt.Println("No es verdadera")
			fallthrough
		case (11 == 14):
			fmt.Println("Tampos es verdadera")
			fallthrough
		case (15 == 15):
			fmt.Println("Si es verdadera")
			fallthrough
		default: // El default en switch de expresion siempre se ejecuta
			fmt.Println("Este es el default")
	}
}
