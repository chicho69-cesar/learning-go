/* 
1. Usa var para DECLARAR tres VARIABLES. Las variables deben tener scope a nivel de
paquete. No asignar VALORES a las variables. Usa los siguientes IDENTIFICADORES
para las variables y asegúrate de que las variables son de los siguientes TIPOS (lo
quiere decir que pueden almacenar VALORES de ese TIPO).
	a. identificador “x” tipo int
	b. identificador “y” tipo string
	c. identificador “z” tipo bool
2. En main
	a. Imprime los valores de cada identificador
	b. El compilador asigna valores a las variables. ¿Cómo son llamados esos valores?
*/

package main 

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
