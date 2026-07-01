/* 
Sigue los siguientes pasos:
	Comienza con este slice
		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	Agrégale el siguiente valor
		52
	Imprime el slice
	En UNA DECLARACIÓN agrega al slice los siguientes valores
		53
		54
		55
	Imprime el slice
	Agregale al Slice los siguientes valores
		y := []int{56, 57, 58, 59, 60}
	print out the slice
*/

package main

import "fmt"

func main() {
	x := []int{ 42, 43, 44, 45, 46, 47, 48, 49, 50, 51 }
	
	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{ 56, 57, 58, 59, 60 }
	x = append(x, y...)
	fmt.Println(x)
}
