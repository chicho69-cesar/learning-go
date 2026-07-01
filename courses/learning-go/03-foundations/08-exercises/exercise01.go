/* 
Escribe un programa que imprima un número en decimal, 
binario, y hexadecimal
*/

package main

import "fmt"

func main() {
	number := 12

	fmt.Printf("El numero en decimal es: %d\n", number)
	fmt.Printf("El numero en binario es: %b\n", number)
	fmt.Printf("El numero en hexadecimal es: %#x\n", number)
}
