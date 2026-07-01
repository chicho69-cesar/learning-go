/* 
En un scope que encierra otros scopes
	Las variables declaradas en el scope externo son accesibles en los scopes
	internos.
Los closures nos ayudan a limitar el scope de las variables
*/

package main

import (
	"fmt"
)

func main() {
	a := incrementor()
	b := incrementor()
	
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}
