/*
Replica el ejercicio anterior pero usa una función
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 10
	y := 10

	ch := generate(x, y)

	for i := 0; i < (x * y); i++ {
		fmt.Println("Valor ", i, ": ", <-ch)
	}

	fmt.Println("Numero de GoRoutines: ", runtime.NumGoroutine())
	fmt.Println("\nFinalizamos...")
}

func generate(x, y int) <-chan int {
	ch := make(chan int)

	for i := 0; i < x; i++ {
		go func() {
			for j := 0; j < y; j++ {
				ch <- j
			}
		}()

		fmt.Println("Numero de GoRoutines: ", runtime.NumGoroutine())
	}

	return ch
}
