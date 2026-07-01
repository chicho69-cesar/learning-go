/* 
Comenzando con este código, extrae los valores del canal usando un ciclo for range
*/

package main

import (
	"fmt"
)

func main() {
	c := generate()
	receive(c)

	fmt.Println("A punto de finalizar.")
}

func generate() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
        fmt.Println(v)
    }
}
