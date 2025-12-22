/* 
Demuestra el uso del idioma coma ok con este código.
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 20
		close(c)
	}()

	v, ok := <- c
	fmt.Println(v, ok)

	v, ok = <- c
	fmt.Println(v, ok)
}
