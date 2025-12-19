/*
Crea una función con el identificador foo que
	Tome un parámetro variable de tipo int
	Pásale un valor de tipo []int a la función (haz el despliegue del []int)
	Retorna la suma de todos los valores de tipo int pasados como argumento.
Crea una func con el identificador bar que
	Tome un parámetro de tipo []int
	Retorne la suma de todos los valores de tipo int pasados.
*/

package main

import "fmt"

func main() {
	nums := []int { 1, 2, 3, 4, 5, 6, 7, 8, 9, 0 }

	sum1 := foo(nums...)
	sum2 := bar(nums)

	fmt.Println("La suma 1 es: ", sum1)
	fmt.Println("La suma 2 es: ", sum2)
}

func foo(nums ...int) int {
	fmt.Println(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

func bar(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
