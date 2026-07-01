/* 
Subyacente a cada slice habrá un arreglo (array). Un slice es realmente una estructura de
datos el cual tiene tres partes:
	1. Un puntero a un arreglo
	2. Longitud (len)
	3. Capacidad (cap)
*/

package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x)
	
	y := append(x, 6, 7, 8) //Un nuevo arreglo subyacente es asignado
	fmt.Println(x)
	fmt.Println(y)
}
