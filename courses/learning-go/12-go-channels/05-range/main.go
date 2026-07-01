/* 
Range deja de leer desde un canal cuando el canal es cerrado.
*/

package main

import (
    "fmt"
)

func main() {
	ch := make(chan int) // creamos el canal de enteros
	
	go func() {
		for i := 0; i < 10; i++ {
            ch <- i // le enviamos los valores
        }

		close(ch) // cerramos el canal para evitar deadLocks!
	}()

	/* Con range podemos iterar sobre los elementos que nos
	envia el canal */
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Finishing")
}
