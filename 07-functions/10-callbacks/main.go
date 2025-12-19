/*
# Pasando una func como un argumento
# La programación funcional no es algo que se recomienda hacer con Go, sin embargo,
es bueno estar al tanto de los callbacks.
# idiomatic go: escribe código que sea claro, simple y legible.
*/

package main

import "fmt"

func main() {
	nums := []int { 1, 2, 3, 4, 5, 6, 7, 8, 9, 0 }

	allNums := sum(nums...)
	fmt.Println("La suma de todos los numeros es:", allNums)

	evenNums := even(sum, nums...) // pasamos la función sum como callback
	fmt.Println("La suma de los numeros pares es:", evenNums)

	oddNums := odd(sum, nums...)
	fmt.Println("La suma de los numeros impares es:", oddNums)

	// podemos definir la función directamente desde la llamada
	specialNums := specialSum(
		func(nums ...int) int {
			total := 0
			for _, v := range nums {
				if v % 3 == 0 {
					total += v
				}
			}
			return total
		}, 
		nums...
	)
	fmt.Println("La suma de los numeros multiplos de 3 es:", specialNums)
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

// recibimos un callback como parámetro de esta función
func even(fn func(nums ...int) int, nums ...int) int {
	var evens []int

	for _, v := range nums {
		if v % 2 == 0 {
			evens = append(evens, v)
		}
	}

	return fn(evens...) // ejecutamos la función o callback que nos llega
}

func odd(fn func(nums ...int) int, nums ...int) int {
	var odds []int

	for _, v := range nums {
		if v % 2 != 0 {
			odds = append(odds, v)
		}
	}

	return fn(odds...)
}

func specialSum(fn func(nums ...int) int, nums ...int) int {
	return fn(nums...)
}
