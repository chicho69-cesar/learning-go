/* 
Crea una variable de tipo string usando un string literal no interpretado 
(raw string literal). Imprímelo
*/

package main 

import "fmt"

func main() {
	var str = `
		Cesar Villalobos Olmos
	`
	
	fmt.Println("El string es: ", str)
}
