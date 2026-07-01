/*
Escribe un programa que:
○ Lance 10 gorutinas
	Cada gorutina agrega 10 números a un canal
○ Extrae los números del canal e imprímelos
*/

package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func(value int) {
			for j := 0; j < 10; j++ {
				ch <- value * j
			}
		}(i)
	}

	for k := 0; k < 100; k++ {
		fmt.Println(k, " -> ", <-ch)
	}

	fmt.Println("Finalizamos...")
}
