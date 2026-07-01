/*
Asigna una función a una variable, luego llama esa función
*/

package main

import "fmt"

var g func()

func main() {
	f := func() {
		for i := 1; i <= 10; i++ {
			fmt.Print(i, " - ")
		}
	}
	f()
	fmt.Printf("\nTipo de f: %T\n", f)
	
	g = f
	g()
	fmt.Printf("\nTipo de g: %T\n", g)
}
