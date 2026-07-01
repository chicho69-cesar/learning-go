/*
Una func que se llama a ella misma
*/

package main

import "fmt"

func main() {
	fmt.Println("1.- La factorial de 5 es:", 5 * 4 * 3 * 2 * 1)
	fmt.Println("2.- La factorial de 5 es:", factorial1(5))
	fmt.Println("3.- La factorial de 5 es:", factorial2(5))
}

// factorial con recursividad
func factorial1(num int) int {
	if num <= 1 {
		return 1
	}
	return num * factorial1(num - 1)
}

// factorial con iteraciones
func factorial2(num int) int {
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}
