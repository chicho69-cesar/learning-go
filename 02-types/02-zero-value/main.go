/* 
El valor 0 en Go es definido por: 
	false para booleanos
	0 para integers
	0.0 para floats
	"" para strings
	nil para: 
		pointers
		funciones
		interfaces
		slices
		channels
		maps
*/

package main

import (
	"fmt"
)

var y int
var x bool
var z float32

func main() {
	fmt.Println(x, y, z)
}
