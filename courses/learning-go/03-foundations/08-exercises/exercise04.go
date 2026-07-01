/* 
Escribe un programa que
	Asignar un int a una variable
	Imprímelo en decimal, binario, y hex
	Has shifts de bits de ese int una posición a la izquierda y asigna eso a una variable
	Imprime esa variable en decimal, binario, y hex
*/

package main 

import "fmt"

func main() {
	var number int = 16
	fmt.Printf("Dec: %d\t\tBin: %b\t\tHex: %#x\n", number, number, number)

	var shifting = number << 1
	fmt.Printf("Dec: %d\t\tBin: %b\t\tHex: %#x\n", shifting, shifting, shifting)
}
