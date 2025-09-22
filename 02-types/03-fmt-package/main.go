/* El paquete fmt se utiliza principalmente para hacer interacciones
con la pantalla y el sistema, como leer datos, mostrar texto, etc. */

package main

import (
	"fmt"
)

var a int
var b string = "Programa"
var c bool
var d bool = true

func main() {
	e := 42
	f := "dice hola mundo" // String interpretado
	g := `El doctor dice que comer vegetales es:
		"saludable".` // Raw String literal

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("Mi", b, f)
	fmt.Println(e)
	fmt.Println(g)
}
