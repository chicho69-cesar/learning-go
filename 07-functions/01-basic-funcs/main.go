/* 
func (receptor) identificador(parámetros) (returns) { código }
Conoce la diferencia entre parámetros y argumentos
	Definimos las funciones con parámetros (si lleva alguno)
	Llamamos las funciones y les pasamos argumentos (si lleva alguno)
-Todo en Go es PASADO POR VALOR
Propósito de las funciones
	Abstraer código
	Reutilización del código
*/

package main

import "fmt"

func main() {
	foo() // mandamos llamar la función
	bar("James") // pasamos argumentos

	s1 := woo("Moneypenny") // función que retorna algo
	fmt.Println(s1)

	x, y := saludar("Eduar", "Tua") // función que regresa varios valores
	fmt.Println(x)
	fmt.Println(y)
}

// creamos una función básica
func foo() {
	fmt.Println("hello from foo")
}

// función que recibe un parámetro
func bar(s string) {
	fmt.Println("Hello,", s)
}

// función que refresca un valor
func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

// función que regresa multiples valores
func saludar(n string, a string) (string, bool) {
	p := fmt.Sprint(n, " ", a, `, dice "Hola."`)
	q := true

	return p, q
}
