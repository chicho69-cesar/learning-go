/* 
Crea constantes CON TIPO y SIN TIPO. Imprime el valor de las mismas
*/

package main

import "fmt"

const number1 = 10
const number2 int = 15

const text1 = "Cesar Villalobos Olmos"
const text2 string = "chicho69-cesar"

const (
	bool1 = true
	bool2 bool = false
)

func main() {
	fmt.Println(number1, number2)
	fmt.Println(text1, text2)
	fmt.Println(bool1, bool2)
}
