package main

import "fmt"

func main() {
	x := 0

	fmt.Println("x befor", &x)
	fmt.Println("x befor", x) // 0

	foo(&x) // pasamos x por referencia

	fmt.Println("x after", &x)
	fmt.Println("x after", x) // 21
}

// recibimos un puntero como parámetro
func foo(y *int) {
	fmt.Println("y befor", y)
	fmt.Println("y befor", *y) // 0 
	
	*y = 21 // modificamos el valor en la posicion de memoria que recibimos
	
	fmt.Println("y after", y)
	fmt.Println("y after", *y) // 21
}
