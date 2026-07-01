/* 
Imprime el rune code point de las letras del alfabeto en mayúsculas tres veces. 
La salida debe ser como esto:
65
	U+0041 'A'
	U+0041 'A'
	U+0041 'A'
66
	U+0042 'B'
	U+0042 'B'
	U+0042 'B'
*/

package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v\n", i)

		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
