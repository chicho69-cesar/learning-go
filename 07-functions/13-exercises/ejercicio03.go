/*
Usa la palabra clave “defer” para mostrar que una función diferida corre después que la función
que la contiene finaliza o retorna.
*/

package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello World!!!")
}

func foo() {
	defer func() {
		fmt.Println("foo DEFERE run")
	}()
	fmt.Println("foo is running")
}
