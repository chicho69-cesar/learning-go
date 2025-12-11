/* 
Podemos usar impresión con formato para imprimir:
	Valor decimal con %d
	Valor hexadecimal con %#x
	Código unipoint %#U
	Una tabulación con \t
	Una línea nueva con \n
*/

package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}
