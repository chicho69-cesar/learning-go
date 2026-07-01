/*
Ejercicio práctico #1
Revisión
	funciones
	Propósito de las funciones
		Código abstracto
		Reusabilidad de código
			DRY - Don’t Repeat Yourself
	func, receptor, identificador, parámetros, retornos
	parámetros vs argumentos
	Variadic func
		Múltiples parámetros “variadic”
		Múltiples argumentos“variadic”
	retornos
		múltiples retornos
		Retornos nombrados - yuck!
	Expresiones func
		Asignando una función a una variable
	callbacks
		Pasando una func a otra func como un argumento
	closure
		Un scope encerrando otro scope
			Variables declaradas en el scope externo son accesibles en el scope
			interno
		Los closure nos ayudan a limitar el scope de las variables
	Recursividad
		Una función que se llama a sí misma
		factorial
Ejercicio práctico
	crea una func con el identificador foo que retorne un int
	crea una func con el identificador bar que retorne un int y un string
	Llama ambas funciones
	Imprime sus resultados
*/

package main

import "fmt"

func main() {
	num1 := foo()
	num2, str1 := bar()

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(str1)
}

func foo() int {
	return 123
}

func bar() (int, string) {
	return 456, "789"
}
